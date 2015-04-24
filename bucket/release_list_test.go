package bucket

import (
	"testing"
)

func TestReleaseListNewReleaseList(t *testing.T) {
	list := NewReleaseList()

	length := list.Len()
	if length != 0 {
		t.Fatalf("Initial length is %d", length)
	}
}

func TestReleaseListAppend(t *testing.T) {
	list := NewReleaseList()
	release, _ := NewRelease("2.0/chromedriver_linux32.zip")

	list = list.Append(release)

	length := list.Len()
	if length != 1 {
		t.Fatalf("Initial length is %d", length)
	}

	r := list.Index(0)
	if r.Version.Major != 2 {
		t.Fatalf("Version.Major missmatch, %d", r.Version.Major)
	}
	if r.Version.Minor != 0 {
		t.Fatalf("Version.Minor missmatch, %d", r.Version.Minor)
	}
	if r.Platform != "linux32" {
		t.Fatalf("Platform missmatch, %s", r.Platform)
	}
}
