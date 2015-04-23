package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a2ikm/chromedriver_helper/chromedriver_helper"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	fmt.Printf("chromedriver_helper v%s\n", Version)

	path, err := chromedriver_helper.Path()
	if err != nil {
		log.Fatal(err)
		return 1
	}

	fmt.Printf("binary path = %s\n", path)
	return 0
}
