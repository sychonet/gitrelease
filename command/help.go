package command

import (
	"fmt"
)

// help0 displays options available at level 0
func help0() {
	fmt.Printf("GitRelease is a tool to automate the process of generating release note on version control systems.\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tgitrelease <command> [arguments]\n\n")
	fmt.Printf("The commands are:\n\n")
	fmt.Printf("\t%-15s\t%-50s\n", "config", "configure the application")
	fmt.Printf("\t%-15s\t%-50s\n", "generate", "generate a new release note for project on vcs")
	fmt.Printf("\t%-15s\t%-50s\n\n", "--version|-v", "display current version")
	fmt.Printf("Type command help to check for valid arguments\n\n")
	fmt.Printf("Source repository:\n\n\thttps://github.com/sychonet/gitrelease\n")
}
