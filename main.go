package main

import (
	"os"

	cmd "github.com/sychonet/gitrelease/command"
)

func main() {
	args := os.Args[1:]
	cmd.Execute(args)
}
