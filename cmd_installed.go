package main

import (
	"fmt"

	"github.com/a2ikm/chromedriver_helper/chromedriver_helper"
)

type CmdInstalled struct{}

func (c *CmdInstalled) Help() string {
	return "Show installed version"
}

func (c *CmdInstalled) Run(args []string) int {
	version, err := chromedriver_helper.InstalledVersion()
	if err != nil {
		fmt.Printf("an error occured: %s\n", err.Error())
		return 1
	}

	fmt.Println(version)
	return 0
}

func (c *CmdInstalled) Synopsis() string {
	return ""
}
