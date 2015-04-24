package bucket

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	KeyRegexp  = regexp.MustCompile("\\A(\\d+)\\.(\\d+)/chromedriver_([a-z0-9]+)\\.zip\\z")
	ParseError = errors.New("parse error")
)

type Release struct {
	Key      string
	Platform string
	Version  Version
}

type Version struct {
	Major int
	Minor int
}

func NewRelease(key string) (*Release, error) {
	major, minor, platform, err := parseKey(key)
	if err != nil {
		return nil, ParseError
	}

	return &Release{
		Key:      key,
		Platform: platform,
		Version: Version{
			Major: major,
			Minor: minor,
		},
	}, nil
}

func parseKey(key string) (int, int, string, error) {
	matches := KeyRegexp.FindStringSubmatch(key)
	if matches == nil {
		return 0, 0, "", ParseError
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	platform := matches[3]

	return major, minor, platform, nil
}

func (r *Release) URL() string {
	return BucketURL + (*r).Key
}
