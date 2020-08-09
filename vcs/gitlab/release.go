package gitlab

import (
	m "github.com/sychonet/gitrelease/model"
	gl "github.com/xanzy/go-gitlab"
)

// newClient creates a new client for interacting with https://www.gitlab.com APIs
func newClient(c m.Config) *gl.Client {
	cl, err := gl.NewClient(c.VCS.Gitlab.AccessToken)
	if err != nil {
		panic("Failed to create client")
	}
	return cl
}

// CreateRelease creates a new release note in gitlab repository
func CreateRelease(args []string, c m.Config) error {
	return nil
}

// GetChangeLog displays change log for gitlab repository
func GetChangeLog(args []string, c m.Config) error {
	return nil
}
