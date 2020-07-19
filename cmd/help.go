package cmd

import (
	"fmt"
)

// help0 displays options available at level 0
func help() {
	fmt.Println("GitRelease is a tool to automate the process of generating release note on version control systems.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgitrelease <command> [arguments]")
	fmt.Println("")
	fmt.Println("The commands are:")
	fmt.Println("")
	fmt.Println("\tconfig\tconfigure the application")
	fmt.Println("\tgenerate\tgenerate a new release note for project on vcs")
	fmt.Println("\t--version|-v\tdisplay current version")
}
