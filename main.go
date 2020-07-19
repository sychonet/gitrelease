package main

import (
	"os"

	"github.com/sychonet/gitrelease/cmd"
)

func main() {
	args := os.Args[1:]
	cmd.Command(args)
}
