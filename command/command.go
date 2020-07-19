package command

import (
	"fmt"
)

// Execute checks for the command given in arguments and takes action correspondingly
func Execute(args []string) {
	if len(args) == 0 {
		help0()
	} else {
		switch args[0] {
		case "help":
			help0()
			break
		case "--version":
			version()
			break
		case "-v":
			version()
			break
		default:
			fmt.Println("Invalid command...")
			fmt.Println("Try help command.")
		}
	}
}
