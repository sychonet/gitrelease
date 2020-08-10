package command

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	m "github.com/sychonet/gitrelease/model"
	"gopkg.in/yaml.v2"
)

// configFile is the filename for gitrelease configuration file
const configFile = ".gitrelease.yaml"

// Configure is used for configuring application
func Configure(cp string) {
CONFIG:
	fmt.Println("#######################")
	fmt.Println("CONFIGURING APPLICATION")
	fmt.Printf("#######################\n\n")
	fmt.Printf("Press enter key to skip\n\n")
	fmt.Printf("Regarding steps for generating personal access tokens refer https://github.com/sychonet/gitrelease/README.md\n\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter personal access token for github.com : ")
	// ght: github token
	ght, gherr := reader.ReadString('\n')
	if gherr != nil {
		fmt.Println("Could not read token from stdin")
		os.Exit(3)
	} else {
		ght = strings.TrimSpace(ght)
	}
	fmt.Printf("Enter personal access token for gitlab.com : ")
	// glt: gitlab token
	glt, glerr := reader.ReadString('\n')
	if glerr != nil {
		fmt.Println("Could not read token from stdin")
		os.Exit(3)
	} else {
		glt = strings.TrimSpace(glt)
	}
	if (len(ght) == 0) && (len(glt) == 0) {
		fmt.Printf("Invalid configuration provided. Please try again...\n\n")
		goto CONFIG
	} else {
		c := &m.Config{}
		c.VCS.Github.AccessToken = ght
		c.VCS.Gitlab.AccessToken = glt

		b, merr := yaml.Marshal(&c)
		if merr != nil {
			fmt.Println("Could not marshal configuration")
			os.Exit(3)
		}
		werr := ioutil.WriteFile(cp, b, 0644)
		if werr != nil {
			fmt.Println("Unable to write confguration to file " + cp)
			os.Exit(3)
		}
		fmt.Printf("\n\nSaved configuration in file " + cp + "\n\n")
	}
}

// GetConfig fetches configuration from file
func GetConfig(cp string) m.Config {
	c := m.Config{}

	// Open configuration file
	f, ferr := os.Open(cp)
	if ferr != nil {
		fmt.Println("Unable to open configuration file")
		os.Exit(3)
	}
	defer f.Close()

	// Init new YAML decoder
	d := yaml.NewDecoder(f)

	// Start YAML decoding from configuration file
	if err := d.Decode(&c); err != nil {
		fmt.Println("Could not decode configuration")
		os.Exit(3)
	}

	return c
}

// GetConfigFileName returns absolute path of the file holding application configuration
func GetConfigFileName() string {
	hd, herr := os.UserHomeDir()
	if herr != nil {
		fmt.Println("Unable to generate path for configuration file")
		os.Exit(3)
	}
	cp := hd + "/" + configFile
	return cp
}
