package cmd

import (
	"fmt"
)

const ver = "0.0.1"

// version displays current version of project
func version() {
	fmt.Println("Current version is " + ver)
}