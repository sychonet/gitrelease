package cmd

import (
	"fmt"
)

// Command checks for the command given in arguments 
func Command(args []string) {
	if len(args) == 0 {
		help()
	} else {
	switch(args[0]) {
		case "help":
			help()
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
