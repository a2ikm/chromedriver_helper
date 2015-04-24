package bucket

import (
	"encoding/xml"
	"io"
	"io/ioutil"
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

	return list, nil
}
