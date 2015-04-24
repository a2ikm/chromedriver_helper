package bucket

import (
	"testing"
)

func TestReleaseNewRelease(t *testing.T) {
	release, err := NewRelease("2.0/chromedriver_linux32.zip")
	if err != nil {
		t.Fatalf("error")
	}

	if release.Version.Major != 2 {
		t.Fatalf("Release.Version.Major missmatch, expected: `2`, actual: `%d`", release.Version.Major)
	}

	if release.Version.Minor != 0 {
		t.Fatalf("Release.Version.Minor missmatch, expected: `0`, actual: `%d`", release.Version.Minor)
	}

	if release.Platform != "linux32" {
		t.Fatalf("Release.Platform missmatch, expected: `linux32`, actual: `%d`", release.Platform)
	}
}

func TestReleaseURL(t *testing.T) {
	release, _ := NewRelease("2.0/chromedriver_linux32.zip")

	expected := "http://chromedriver.storage.googleapis.com/2.0/chromedriver_linux32.zip"
	if release.URL() != expected {
		t.Fatalf("Release.URL() missmatch, expected: `%s`, actual: `%s`", expected, release.URL())
	}
}
