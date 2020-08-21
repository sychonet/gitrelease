package command

import (
	"fmt"
)

// help displays options available
func help() {
	fmt.Printf("GitRelease is a tool to automate the process of generating release note on version control systems.\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tgitrelease <command> [arguments]\n\n")
	fmt.Printf("The commands are:\n\n")
	fmt.Printf("\t%-15s\t%-50s\n", "configure", "configure the application")
	fmt.Printf("\t%-15s\t%-50s\n", "create", "generate a new draft release note for project on vcs")
	fmt.Printf("\t%-15s\t%-50s\n", "view", "list change log for project on vcs")
	fmt.Printf("\t%-15s\t%-50s\n\n", "version", "display current version")
	fmt.Printf("Type command help to check for valid arguments if available\n\n")
	fmt.Printf("Source repository:\n\n\thttps://github.com/sychonet/gitrelease\n")
}

// createHelp displays options available for create command
func createHelp() {
	fmt.Printf("Create command is used to generate release note.\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tgitrelease create --vcs <vcs> --owner <owner> --repo <repository> --latest <latestRelease> --previous <previousRelease> --custom\n\n")
	fmt.Printf("The arguments are:\n\n")
	fmt.Printf("\t%-15s\t%-50s\n", "--owner|-o", "owner of repository")
	fmt.Printf("\t%-15s\t%-50s\n", "--vcs|-v", "version control system (github)")
	fmt.Printf("\t%-15s\t%-50s\n", "--repo|-r", "name of repository")
	fmt.Printf("\t%-15s\t%-50s\n", "--latest|-l", "latest release of project i.e. latest tag number")
	fmt.Printf("\t%-15s\t%-50s\n", "--previous|-p", "previous release of project i.e. previous release tag number (optional)")
	fmt.Printf("\t%-15s\t%-50s\n", "--custom|-c", "add custom information in release note (optional)")
}

// viewHelp displays options available for view command
func viewHelp() {
	fmt.Printf("View command is used to display change log.\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tgitrelease view --vcs <vcs> --owner <owner> --repo <repository> --latest <latestRelease> --previous <previousRelease>\n\n")
	fmt.Printf("The arguments are:\n\n")
	fmt.Printf("\t%-15s\t%-50s\n", "--owner|-o", "owner of repository")
	fmt.Printf("\t%-15s\t%-50s\n", "--vcs|-v", "version control system (github)")
	fmt.Printf("\t%-15s\t%-50s\n", "--repo|-r", "name of repository")
	fmt.Printf("\t%-15s\t%-50s\n", "--latest|-l", "latest release of project i.e. latest tag number")
	fmt.Printf("\t%-15s\t%-50s\n", "--previous|-p", "previous release of project i.e. previous release tag number (optional)")
}
