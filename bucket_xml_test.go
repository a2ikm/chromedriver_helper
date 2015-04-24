package main

import (
	"testing"
)

func TestParseBucketXML(t *testing.T) {
	data := []byte(`
  <?xml version='1.0' encoding='UTF-8'?>
  <ListBucketResult xmlns="http://doc.s3.amazonaws.com/2006-03-01">
    <Name>chromedriver</Name>
    <Prefix/>
    <Marker/>
    <IsTruncated>false</IsTruncated>
    <Contents>
      <Key>2.0/chromedriver_linux32.zip</Key>
      <Generation>1380149859530000</Generation>
      <MetaGeneration>2</MetaGeneration>
      <LastModified>2013-09-25T22:57:39.349Z</LastModified>
      <ETag>"c0d96102715c4916b872f91f5bf9b12c"</ETag>
      <Size>7262134</Size>
      <Owner/>
    </Contents>
  </ListBucketResult>
  `)

	result, err := parseBucketXML(data)
	if err != nil {
		t.Fatalf("error : %s", err)
	}

	if result.Name != "chromedriver" {
		t.Fatalf("Name missmatch, expected: `chromedriver`, actual: `%s`", result.Name)
	}

	if result.Prefix != "" {
		t.Fatalf("Prefix missmatch, expected: ``, actual: `%s`", result.Prefix)
	}

	if result.Marker != "" {
		t.Fatalf("Marker missmatch, expected: ``, actual: `%s`", result.Marker)
	}

	if result.IsTruncated != "false" {
		t.Fatalf("IsTruncated missmatch, expected: `false`, actual: `%s", result.IsTruncated)
	}

	contents := result.Contents[0]

	if contents.Key != "2.0/chromedriver_linux32.zip" {
		t.Fatalf("Contents.Key missmatch, expected: `2.0/chromedriver_linux32.zip`, actual: `%s`", contents.Key)
	}

	if contents.Generation != 1380149859530000 {
		t.Fatalf("Contents.Generation missmatch, expected: `1380149859530000`, actual: `%d`", contents.Generation)
	}

	if contents.MetaGeneration != 2 {
		t.Fatalf("Contents.MetaGeneration missmatch, expected: `2`, actual: `%d`", contents.MetaGeneration)
	}

	if contents.LastModified != "2013-09-25T22:57:39.349Z" {
		t.Fatalf("Contents.LastModified missmatch, expected: `2013-09-25T22:57:39.349Z`, actual: `%s`", contents.LastModified)
	}

	if contents.ETag != "\"c0d96102715c4916b872f91f5bf9b12c\"" {
		t.Fatalf("Contents.ETag missmatch, expected: `\"c0d96102715c4916b872f91f5bf9b12c\"`, actual: `%s`", contents.ETag)
	}

	if contents.Size != 7262134 {
		t.Fatalf("Contents.Size missmatch, expected: `7262134`, actual: `%d`", contents.Size)
	}
}

func TestContentsConvertToRelease(t *testing.T) {
	contents := Contents{
		Key:            "2.0/chromedriver_linux32.zip",
		Generation:     1380149859530000,
		MetaGeneration: 2,
		LastModified:   "2013-09-25T22:57:39.349Z",
		ETag:           "\"c0d96102715c4916b872f91f5bf9b12c\"",
		Size:           7262134,
		Owner:          "",
	}

	release, err := contents.ConvertToRelease()
	if err != nil {
		t.Fatalf("error")
	}

	if release.Key != "2.0/chromedriver_linux32.zip" {
		t.Fatalf("Release.Key missmatch, expected: `2.0/chromedriver_linux32.zip`, actual: `%s`", release.Key)
	}

	if release.Platform != "linux32" {
		t.Fatalf("Release.Platform missmatch, expected: `linux32`, actual: `%s`", release.Platform)
	}

	if release.Version.Major != 2 {
		t.Fatalf("Release.Version.Major missmatch, expected: `2`, actual: `%d`", release.Version.Major)
	}

	if release.Version.Minor != 0 {
		t.Fatalf("Release.Version.Minor missmatch, expected: `0`, actual: `%d`", release.Version.Minor)
	}
}
