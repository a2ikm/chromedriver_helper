package chromedriver_helper

import (
	"testing"
)

func TestParseVersion(t *testing.T) {
	out := "ChromeDriver 2.13.307650 (feffe1dd547ee7b5c16d38784cd0cd679dfd7850)"
	version, err := parseVersion(out)

	if err != nil {
		t.Fatalf("error, %s", err.Error())
	}

	if version != "2.13" {
		t.Fatalf("missmatch %s", version)
	}
}
