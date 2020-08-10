package gitlab

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	m "github.com/sychonet/gitrelease/model"
	u "github.com/sychonet/gitrelease/util"
	gl "github.com/xanzy/go-gitlab"
)

// newClient creates a new client for interacting with https://www.gitlab.com APIs
func newClient(c m.Config) *gl.Client {
	cl, err := gl.NewClient(c.VCS.Gitlab.AccessToken)
	if err != nil {
		fmt.Println("Failed to create client")
		os.Exit(3)
	}
	return cl
}

// CreateRelease creates a new release note in gitlab repository
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

	pid := o + "/" + rn

	// Fetch previous tag
	emsg = "Repository owner missing in arguments"
	p, err := u.GetOptionValue(args, emsg, "-p", "--previous")
	if err == nil {
		lrt = p
	} else {
		// rel : repository release
		rel, _, err := client.Releases.ListReleases(pid, nil, nil)
		if err != nil || len(rel) == 0 {
			fmt.Println("Unable to fetch latest release for repository")
			os.Exit(3)
		}

		// lrt: latest release tag
		lrt = rel[0].TagName

		fmt.Println("Tag of latest release for project is " + lrt)
	}

	cmpo := &gl.CompareOptions{From: &lrt, To: &lt}
	cmp, _, err := client.Repositories.Compare(pid, cmpo, nil)
	if err != nil {
		fmt.Println("Could not fetch merge requests between tags")
		os.Exit(3)
	}

	// rp: regex pattern for fetching merge request number
	rp := regexp.MustCompile(`![0-9]*`)

	// rln: release note
	var rln string

	if len(cmp.Commits) > 0 {
		for i := 0; i < len(cmp.Commits); i++ {
			if strings.Contains(cmp.Commits[i].Message, "See merge request") {
				// prn: pull request number
				prn := rp.FindString(cmp.Commits[i].Message)
				// prt: pull request title
				prt := strings.Split(cmp.Commits[i].Message, "\n")[2]
				rln += "- " + prn + " " + prt + "\n"
			}
		}
	} else {
		fmt.Println("No change log available")
		os.Exit(3)
	}

	// lrn: latest release name
	lrn := "Releasing " + lt

	// cro: create release options
	cro := &gl.CreateReleaseOptions{Name:&lrn, TagName:&lt, Description:&rln}
	_, _, err = client.Releases.CreateRelease(pid, cro, nil)

	if err != nil {
		fmt.Println("Unable to create a new release note")
		os.Exit(3)
	} else {
		fmt.Println("Release note successfully created for tag " + lt + " in repository " + o + "/" + rn)
	}

	return nil
}

// GetChangeLog displays change log for gitlab repository
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

	pid := o + "/" + rn

	// Fetch previous tag
	emsg = "Repository owner missing in arguments"
	p, err := u.GetOptionValue(args, emsg, "-p", "--previous")
	if err == nil {
		lrt = p
	} else {
		// arel : array of repository releases
		arel, _, err := client.Releases.ListReleases(pid, nil, nil)
		if err != nil || len(arel) == 0 {
			fmt.Println(err)
			fmt.Println("Unable to fetch latest release for repository")
			os.Exit(3)
		}

		// lrt: latest release tag
		lrt = arel[0].TagName

		fmt.Println("Tag of latest release for project is " + lrt)
	}

	cmpo := &gl.CompareOptions{From: &lrt, To: &lt}
	cmp, _, err := client.Repositories.Compare(pid, cmpo, nil)
	
	if err != nil {
		fmt.Println("Could not fetch merge requests between tags")
		os.Exit(3)
	}

	if len(cmp.Commits) == 0 {
		fmt.Println("No change log available")
		os.Exit(3)
	}

	// rp: regex pattern for fetching merge request number
	rp := regexp.MustCompile(`![0-9]*`)

	// rln: release note
	var rln string

	for i := 0; i < len(cmp.Commits); i++ {
		if strings.Contains(cmp.Commits[i].Message, "See merge request") {
			// prn: pull request number
			prn := rp.FindString(cmp.Commits[i].Message)
			// prt: pull request title
			prt := strings.Split(cmp.Commits[i].Message, "\n")[2]
			rln += "- " + prn + " " + prt + "\n"
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
