package github

import (
	"context"
	"fmt"
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
	var rn, lt, o, emsg string
	// err: error
	var err error

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

	// client: github client
	client := newClient(c)

	// rel : repository release
	rel, _, err := client.Repositories.GetLatestRelease(context.Background(), o, rn)
	if err != nil {
		panic("Unable to fetch latest release for repository")
	}

	// lrt: latest release tag
	lrt := rel.GetTagName()

	fmt.Println("Tag of latest release for project is " + lrt)

	// cc: commits comparison
	cc, _, err := client.Repositories.CompareCommits(context.Background(), o, rn, lrt, lt)
	if err != nil {
		panic("Unable to fetch comparison between tags for repository")
	}

	// rp: regex pattern for fetching merge request number
	rp := regexp.MustCompile(`#[0-9]*`)

	// rln: release note
	rln := ""

	for i := 0; i < len(cc.Commits); i++ {
		if strings.Contains(cc.Commits[i].Commit.GetMessage(), "Merge pull request") {
			// prn: pull request number
			prn := rp.FindString(cc.Commits[i].Commit.GetMessage())
			// prt: pull request title
			prt := strings.Split(cc.Commits[i].Commit.GetMessage(), "\n")[2]
			rln += "- " + prn + " " + prt + "<br/>"
		}
	}

	fmt.Println("Release note for lastest tag:")
	fmt.Println(rln)

	// rr: repository release
	lrn := "Releasing " + lrt
	// dr: draft
	dr := true
	rr := &gh.RepositoryRelease{TagName: &lrt, Name: &lrn, Draft: &dr, Body: &rln}

	rr, _, err = client.Repositories.CreateRelease(context.Background(), o, rn, rr)

	if err != nil {
		panic("Unable to create a new release note")
	} else {
		fmt.Println("Draft release note successfully created for tag " + lt + " in repository " + o + "/" + rn)
	}

	return nil
}
