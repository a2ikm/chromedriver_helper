package bucket

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	BucketURL = "http://chromedriver.storage.googleapis.com/"
)

type ListBucketResult struct {
	Name        string
	Prefix      string
	Marker      string
	IsTruncated string
	Contents    []Contents
}

type Contents struct {
	Key            string
	Generation     int
	MetaGeneration int
	LastModified   string
	ETag           string
	Size           int
	Owner          string
}

func DownloadReleaseList() (*ReleaseList, error) {
	res, err := http.Get(BucketURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ParseXML(res.Body)
}

func ParseXML(in io.Reader) (*ReleaseList, error) {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	result := ListBucketResult{}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	list := NewReleaseList()
	var release *Release
	for _, c := range result.Contents {
		release, err = NewRelease(c.Key)
		if err != nil {
			if err == ParseError {
				continue
			} else {
				return nil, err
			}
		}
		list = list.Append(release)
	}

	list.Sort()
	return list, nil
}

func LatestReleaseForPlatform(platform string) (*Release, error) {
	list, err := DownloadReleaseList()
	if err != nil {
		return nil, err
	}

	return list.FilterByPlatform(platform).Latest(), nil
}
