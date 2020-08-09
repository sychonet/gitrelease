package command

import (
	"fmt"

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

// create creates a new draft release note on vcs
func create(args []string, c m.Config) {
	p := u.SliceIndex(len(args), func(i int) bool { return args[i] == "-vc" })
	if p < 0 {
		p = u.SliceIndex(len(args), func(i int) bool { return args[i] == "--vcs" })
	}
	if p >= 0 {
		switch args[p+1] {
		case "github":
			err := gh.CreateRelease(args, c)
			if err != nil {
				fmt.Println(err.Error())
				panic("Unable to create release note in github repository")
			}
			break
		case "gitlab":
			err := gl.CreateRelease(args, c)
			if err != nil {
				fmt.Println(err.Error())
				panic("Unable to create release note in gitlab repository")
			}
			break
		case "default":
			fmt.Println("Invalid argument value")
			createHelp()
		}
	}
}

// view displays change log
func view(args []string, c m.Config) {
	p := u.SliceIndex(len(args), func(i int) bool { return args[i] == "-vc" })
	if p < 0 {
		p = u.SliceIndex(len(args), func(i int) bool { return args[i] == "--vcs" })
	}
	if p >= 0 {
		switch args[p+1] {
		case "github":
			err := gh.GetChangeLog(args, c)
			if err != nil {
				fmt.Println(err.Error())
				panic("Unable to get change log of github repository")
			}
			break
		case "gitlab":
			err := gl.GetChangeLog(args, c)
			if err != nil {
				fmt.Println(err.Error())
				panic("Unable to get change log of gitlab repository")
			}
			break
		case "default":
			fmt.Println("Invalid argument value")
			createHelp()
		}
	}
}
