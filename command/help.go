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
	fmt.Printf("\t%-15s\t%-50s\n", "config", "configure the application")
	fmt.Printf("\t%-15s\t%-50s\n", "create", "generate a new release note for project on vcs")
	fmt.Printf("\t%-15s\t%-50s\n\n", "--version|-v", "display current version")
	fmt.Printf("Type command help to check for valid arguments\n\n")
	fmt.Printf("Source repository:\n\n\thttps://github.com/sychonet/gitrelease\n")
}

// createHelp displays options available for create command
func createHelp() {
	fmt.Printf("Create command is used to generate release note.\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tgitrelease create --vcs <vcs> --owner <owner> --repo <repository> --latest <latestRelease> --previous <previousRelease>\n\n")
	fmt.Printf("The arguments are:\n\n")
	fmt.Printf("\t%-15s\t%-50s\n", "--owner|-o", "owner of repository")
	fmt.Printf("\t%-15s\t%-50s\n", "--vcs|-vc", "version control system (github)")
	fmt.Printf("\t%-15s\t%-50s\n", "--repo|-r", "name of repository")
	fmt.Printf("\t%-15s\t%-50s\n", "--latest|-l", "latest release of project i.e. lastest tag number")
	fmt.Printf("\t%-15s\t%-50s\n", "--previous|-p", "previous release of project i.e. previous release tag number (optional)")
}
