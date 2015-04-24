package chromedriver_helper

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

var (
	ParseError    = errors.New("parse error")
	VersionRegexp = regexp.MustCompile("\\AChromeDriver (\\d+\\.\\d+)")
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

func InstalledVersion() (string, error) {
	path, err := Path()
	if err != nil {
		return "", err
	}

	_, err = os.Stat(path)
	if err != nil {
		return "", err
	}

	cmd := fmt.Sprintf("%s --version", path)
	out, err := exec.Command(cmd).Output()
	if err != nil {
		return "", err
	}

	version, err := parseVersion(string(out))
	if err != nil {
		return "", err
	}

	return version, nil
}

func parseVersion(s string) (string, error) {
	matches := VersionRegexp.FindStringSubmatch(s)
	if matches == nil {
		return "", ParseError
	}

	return matches[1], nil
}
