package github

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	gh "github.com/google/go-github/v32/github"
	m "github.com/sychonet/gitrelease/model"
	u "github.com/sychonet/gitrelease/util"

	"golang.org/x/oauth2"
)

// newClient creates a new client for interacting with https://www.github.com APIs
func newClient(c m.Config) *gh.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.VCS.Github.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := gh.NewClient(tc)

	return client
}

// CreateRelease creates a new release note in github repository
func CreateRelease(args []string, c m.Config) error {
	// rn: repository name
	// lt: latest tag
	// o: owner
	// emsg: error message
	var rn, lt, o, emsg, lrt string
	// err: error
	var err error

	// client: github client
	client := newClient(c)

	// Fetch repository name
	emsg = "Repository name missing in arguments"
	rn, err = u.GetOptionValue(args, emsg, "-r", "--repo")
	if err != nil {
		return err
	}

	// Fetch latest release tag
	emsg = "Latest release tag missing in arguments"
	lt, err = u.GetOptionValue(args, emsg, "-l", "--latest")
	if err != nil {
		return err
	}

	// Fetch owner name
	emsg = "Repository owner missing in arguments"
	o, err = u.GetOptionValue(args, emsg, "-o", "--owner")
	if err != nil {
		return err
	}

	// Fetch previous tag
	emsg = "Repository owner missing in arguments"
	p, err := u.GetOptionValue(args, emsg, "-p", "--previous")
	if err == nil {
		lrt = p
	} else {
		// rel : repository release
		rel, _, err := client.Repositories.GetLatestRelease(context.Background(), o, rn)
		if err != nil {
			fmt.Println("Unable to fetch latest release for repository")
			os.Exit(3)
		}

		// lrt: latest release tag
		lrt = rel.GetTagName()

		fmt.Println("Tag of latest release for project is " + lrt)
	}

	// cc: commits comparison
	cc, _, err := client.Repositories.CompareCommits(context.Background(), o, rn, lrt, lt)

	if err != nil {
		fmt.Println("Unable to fetch comparison between tags for repository")
		os.Exit(3)
	}

	if len(cc.Commits) == 0 {
		fmt.Println("No change log available")
		os.Exit(3)
	}

	// rp: regex pattern for fetching merge request number
	rp := regexp.MustCompile(`#[0-9]*`)

	// rln: release note
	var rln string

	for i := 0; i < len(cc.Commits); i++ {
		if strings.Contains(cc.Commits[i].Commit.GetMessage(), "Merge pull request") {
			// prn: pull request number
			prn := rp.FindString(cc.Commits[i].Commit.GetMessage())
			// prt: pull request title
			prt := strings.Split(cc.Commits[i].Commit.GetMessage(), "\n")[2]
			rln += prn + " " + prt + "\n"
		}
	}

	if len(rln) == 0 {
		fmt.Println("No change log available")
		os.Exit(3)
	}

	fmt.Println("Release note for lastest tag:")
	fmt.Println(rln)

	// lrn: latest release name
	lrn := "Releasing " + lt
	// dr: draft
	dr := true

	tmp := u.SliceIndex(len(args), func(i int) bool { return args[i] == "-c" })
	if tmp < 0 {
		tmp = u.SliceIndex(len(args), func(i int) bool { return args[i] == "--custom" })
	}

	if tmp >= 0 {
		reader := bufio.NewReader(os.Stdin)

		// Give option to edit release name using an editor
		fmt.Printf("Want to customize release name (y/n)? ")
		// opn: option to change name
		opn, opnerr := reader.ReadString('\n')
		if opnerr != nil {
			fmt.Println("Could not read input from stdin")
			os.Exit(3)
		}
		opn = strings.TrimSpace(opn)
		if (opn == "y") || (opn == "Y") {
			lrnBytes, err := u.CaptureInputFromEditor(lrn, u.GetPreferredEditorFromEnvironment)
			if len(lrnBytes) == 0 || err != nil {
				fmt.Println("Unable to edit release name. Using default release name instead.")
			} else {
				lrn = string(lrnBytes)
			}
		}

		// Give option to edit release description using an editor
		fmt.Printf("Want to customize release description (y/n)? ")
		// opd: option to change description
		opd, opderr := reader.ReadString('\n')
		if opderr != nil {
			fmt.Println("Could not read input from stdin")
			os.Exit(3)
		}
		opd = strings.TrimSpace(opd)
		if (opd == "y") || (opd == "Y") {
			rlnBytes, err := u.CaptureInputFromEditor(rln, u.GetPreferredEditorFromEnvironment)
			if len(rlnBytes) == 0 || err != nil {
				fmt.Println("Unable to edit release description. Using default release description instead.")
			} else {
				rln = string(rlnBytes)
			}
		}
	}

	// rr: repository release
	rr := &gh.RepositoryRelease{TagName: &lt, Name: &lrn, Draft: &dr, Body: &rln}
	rr, _, err = client.Repositories.CreateRelease(context.Background(), o, rn, rr)

	if err != nil {
		fmt.Println("Unable to create a new release note")
		os.Exit(3)
	} else {
		fmt.Println("Draft release note successfully created for tag " + lt + " in repository " + o + "/" + rn)
	}

	return nil
}

// GetChangeLog displays change log for github repository
func GetChangeLog(args []string, c m.Config) error {
	// rn: repository name
	// lt: latest tag
	// o: owner
	// emsg: error message
	var rn, lt, o, emsg, lrt string
	// err: error
	var err error

	// client: github client
	client := newClient(c)

	// Fetch repository name
	emsg = "Repository name missing in arguments"
	rn, err = u.GetOptionValue(args, emsg, "-r", "--repo")
	if err != nil {
		return err
	}

	// Fetch latest release tag
	emsg = "Latest release tag missing in arguments"
	lt, err = u.GetOptionValue(args, emsg, "-l", "--latest")
	if err != nil {
		return err
	}

	// Fetch owner name
	emsg = "Repository owner missing in arguments"
	o, err = u.GetOptionValue(args, emsg, "-o", "--owner")
	if err != nil {
		return err
	}

	// Fetch previous tag
	emsg = "Repository owner missing in arguments"
	p, err := u.GetOptionValue(args, emsg, "-p", "--previous")
	if err == nil {
		lrt = p
	} else {
		// rel : repository release
		rel, _, err := client.Repositories.GetLatestRelease(context.Background(), o, rn)
		if err != nil {
			fmt.Println("Unable to fetch latest release for repository")
			os.Exit(3)
		}

		// lrt: latest release tag
		lrt = rel.GetTagName()

		fmt.Println("Tag of latest release for project is " + lrt)
	}

	// cc: commits comparison
	cc, _, err := client.Repositories.CompareCommits(context.Background(), o, rn, lrt, lt)

	if err != nil {
		fmt.Println("Unable to fetch comparison between tags for repository")
		os.Exit(3)
	}

	if len(cc.Commits) == 0 {
		fmt.Println("No change log available")
		os.Exit(3)
	}

	// rp: regex pattern for fetching merge request number
	rp := regexp.MustCompile(`#[0-9]*`)

	// rln: release note
	var rln string

	for i := 0; i < len(cc.Commits); i++ {
		if strings.Contains(cc.Commits[i].Commit.GetMessage(), "Merge pull request") {
			// prn: pull request number
			prn := rp.FindString(cc.Commits[i].Commit.GetMessage())
			// prt: pull request title
			prt := strings.Split(cc.Commits[i].Commit.GetMessage(), "\n")[2]
			rln += prn + " " + prt + "\n"
		}
	}

	if len(rln) == 0 {
		fmt.Println("No change log available")
		os.Exit(3)
	}

	fmt.Printf("Change log for repository:\n\n")
	fmt.Println(rln)

	return nil
}
