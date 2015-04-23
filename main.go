package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	fmt.Printf("chromedriver_helper v%s\n", Version)
	return 0
}
