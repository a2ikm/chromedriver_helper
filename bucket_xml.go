package main

import (
	"encoding/xml"
	"errors"
	"strings"

	"github.com/a2ikm/chromedriver_helper/bucket"
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

func parseBucketXML(data []byte) (ListBucketResult, error) {
	list := ListBucketResult{}
	err := xml.Unmarshal(data, &list)
	if err != nil {
		return ListBucketResult{}, err
	}

	return list, nil
}

func (c *Contents) ConvertToRelease() (bucket.Release, error) {
	comp := strings.Split(c.Key, "/")
	if len(comp) != 2 {
		return bucket.Release{}, errors.New("Not Chromedriver")
	}

	filename := comp[1]
	if !strings.HasPrefix(filename, "chromedriver") {
		return bucket.Release{}, errors.New("Not Chromedriver")
	}

	release, err := bucket.NewRelease(c.Key)
	if err != nil {
		return bucket.Release{}, err
	}

	return *release, nil
}
