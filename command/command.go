package command

import (
	"fmt"
	"os"

	m "github.com/sychonet/gitrelease/model"
	u "github.com/sychonet/gitrelease/util"
	gh "github.com/sychonet/gitrelease/vcs/github"
	gl "github.com/sychonet/gitrelease/vcs/gitlab"
)

// Execute checks for the command given in arguments and takes action correspondingly
func Execute(args []string, c m.Config) {
	if len(args) == 0 {
		help()
	} else {
		switch args[0] {
		case "configure":
			cp := GetConfigFileName()
			Configure(cp)
			break
		case "create":
			create(args[1:], c)
			break
		case "view":
			view(args[1:], c)
			break
		case "help":
			help()
			break
		case "version":
			version()
			break
		default:
			fmt.Println("Invalid command...")
			fmt.Println("Try help command.")
		}
	}
}

// create creates a new draft release note on vcs
func create(args []string, c m.Config) {
	p := u.SliceIndex(len(args), func(i int) bool { return args[i] == "-v" })
	if p < 0 {
		p = u.SliceIndex(len(args), func(i int) bool { return args[i] == "--vcs" })
	}
	if p >= 0 {
		switch args[p+1] {
		case "github":
			err := gh.CreateRelease(args, c)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Unable to create release note in github repository")
				os.Exit(3)
			}
			break
		case "gitlab":
			err := gl.CreateRelease(args, c)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Unable to create release note in gitlab repository")
				os.Exit(3)
			}
			break
		case "default":
			fmt.Println("Invalid argument value")
			createHelp()
		}
	} else {
		createHelp()
	}
}

// view displays change log
func view(args []string, c m.Config) {
	p := u.SliceIndex(len(args), func(i int) bool { return args[i] == "-v" })
	if p < 0 {
		p = u.SliceIndex(len(args), func(i int) bool { return args[i] == "--vcs" })
	}
	if p >= 0 {
		switch args[p+1] {
		case "github":
			err := gh.GetChangeLog(args, c)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Unable to get change log of github repository")
				os.Exit(3)
			}
			break
		case "gitlab":
			err := gl.GetChangeLog(args, c)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Unable to get change log of gitlab repository")
				os.Exit(3)
			}
			break
		case "default":
			fmt.Println("Invalid argument value")
			viewHelp()
		}
	} else {
		viewHelp()
	}
}
