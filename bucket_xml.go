package main

import (
	"encoding/xml"
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
