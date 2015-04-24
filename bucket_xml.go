package main

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

type Contents struct {
	Key            string
	Generation     int
	MetaGeneration int
	LastModified   string
	ETag           string
	Size           int
	Owner          string
}

type ListBucketResult struct {
	Name        string
	Prefix      string
	Marker      string
	IsTruncated string
	Contents    []Contents
}

type Release struct {
	Key      string
	Platform string
	Version  ReleaseVersion
}

type ReleaseVersion struct {
	Major int
	Minor int
}

func parseBucketXML(data []byte) (ListBucketResult, error) {
	list := ListBucketResult{}
	err := xml.Unmarshal(data, &list)
	if err != nil {
		return ListBucketResult{}, err
	}

	return list, nil
}

func (c *Contents) ConvertToRelease() (Release, error) {
	comp := strings.Split(c.Key, "/")
	if len(comp) != 2 {
		return Release{}, errors.New("Not Chromedriver")
	}

	version := comp[0]
	filename := comp[1]

	if !strings.HasPrefix(filename, "chromedriver") {
		return Release{}, errors.New("Not Chromedriver")
	}

	major, minor, err := c.parseVersion(version)
	if err != nil {
		return Release{}, err
	}

	platform := c.parsePlatform(filename)

	release := Release{
		Key:      c.Key,
		Platform: platform,
		Version: ReleaseVersion{
			Major: major,
			Minor: minor,
		},
	}

	return release, nil
}

func (c *Contents) parseVersion(version string) (int, int, error) {
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

func (c *Contents) parsePlatform(filename string) string {
	comp1 := strings.Split(filename, ".")
	comp2 := strings.Split(comp1[0], "_")
	return comp2[1]
}
