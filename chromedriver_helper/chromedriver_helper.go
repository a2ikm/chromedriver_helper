package chromedriver_helper

import (
	"errors"
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

func BinaryName() string {
	if runtime.GOOS == "windows" {
		return "chromedriver.exe"
	}

	return "chromedriver"
}

func BinaryPath() (string, error) {
	dir, err := InstallDir()
	if err != nil {
		return "", err
	}

	name := BinaryName()
	return path.Join(dir, name), nil
}

func InstallDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".chromedriver-helper"), nil
}

func MakeInstallDir() error {
	dir, err := InstallDir()
	if err != nil {
		return err
	}

	return os.MkdirAll(dir, 0755)
}

func InstalledVersion() (string, error) {
	path, err := BinaryPath()
	if err != nil {
		return "", err
	}

	_, err = os.Stat(path)
	if err != nil {
		return "", err
	}

	out, err := exec.Command(path, "--version").Output()
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
