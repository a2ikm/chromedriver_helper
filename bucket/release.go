package bucket

import (
	"strconv"
	"strings"
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
	comp := strings.Split(key, "/")

	major, minor, err := parseVersion(comp[0])
	if err != nil {
		return &Release{}, err
	}

	platform := parsePlatform(comp[1])

	return &Release{
		Key:      key,
		Platform: platform,
		Version: Version{
			Major: major,
			Minor: minor,
		},
	}, nil
}

func parseVersion(version string) (int, int, error) {
	comp := strings.Split(version, ".")
	major, err := strconv.Atoi(comp[0])
	if err != nil {
		return 0, 0, err
	}
	minor, err := strconv.Atoi(comp[1])
	if err != nil {
		return 0, 0, err
	}
	return major, minor, nil
}

func parsePlatform(filename string) string {
	comp1 := strings.Split(filename, ".")
	comp2 := strings.Split(comp1[0], "_")
	return comp2[1]
}
