package main

import (
	"fmt"
	"log"

	"github.com/a2ikm/chromedriver_helper/chromedriver_helper"
)

type CmdVersion struct{}

func (c *CmdVersion) Help() string {
	return "Show version"
}

func (c *CmdVersion) Run(args []string) int {
	c.printVersion()
	c.printChromedriverPath()
	return 0
}

func (c *CmdVersion) Synopsis() string {
	return ""
}

func (c *CmdVersion) printVersion() {
	fmt.Printf("chromedriver_helper v%s\n", chromedriver_helper.Version)
}

func (c *CmdVersion) printChromedriverPath() {
	path, err := chromedriver_helper.Path()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("binary path = %s\n", path)
}
