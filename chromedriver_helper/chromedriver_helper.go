package chromedriver_helper

import (
	"path"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

func Name() string {
	if runtime.GOOS == "windows" {
		return "chromedriver.exe"
	}

	return "chromedriver"
}

func Dir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".chromedriver-helper"), nil
}

func Path() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, Name()), nil
}
