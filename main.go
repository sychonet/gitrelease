package main

import (
	"os"
	cmd "github.com/sychonet/gitrelease/command"
)

func main() {
	cp := cmd.GetConfigFileName()
	c := cmd.GetConfig(cp)
	args := os.Args[1:]
	cmd.Execute(args, c)
}

func init() {
	cp := cmd.GetConfigFileName()
	if _, err := os.Stat(cp); os.IsNotExist(err) {
		// Configure applicaton if not already configured
		cmd.Configure(cp)
	}
}
