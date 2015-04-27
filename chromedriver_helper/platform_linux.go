package chromedriver_helper

import (
	"syscall"
)

func Platform() (string, error) {
	m, err := machine()
	if err != nil {
		return "", err
	}

	if m == "x86_64" || m == "amd64" {
		return "linux64", nil
	} else {
		return "linux32", nil
	}
}

func machine() (string, error) {
	u := syscall.Utsname{}
	err := syscall.Uname(&u)
	if err != nil {
		return "", err
	}

	var m string
	for _, val := range u.Machine {
		m += string(int(val))
	}

	return m, nil
}
