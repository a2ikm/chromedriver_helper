package bucket

import (
	"bytes"
	"testing"
)

func TestBucketParseXML(t *testing.T) {
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

	in := bytes.NewReader(data)
	list, err := ParseXML(in)
	if err != nil {
		t.Fatalf("fatal")
	}

	length := list.Len()
	if length != 1 {
		t.Fatalf("length missmatch, %d", length)
	}

	r := list.Index(0)
	if r.Key != "2.0/chromedriver_linux32.zip" {
		t.Fatalf("Release.Key missmatch, `%s`", r.Key)
	}
}

func TestBucketParseXMLNonChromedriver(t *testing.T) {
	data := []byte(`
  <?xml version='1.0' encoding='UTF-8'?>
  <ListBucketResult xmlns="http://doc.s3.amazonaws.com/2006-03-01">
    <Name>chromedriver</Name>
    <Prefix/>
    <Marker/>
    <IsTruncated>false</IsTruncated>
    <Contents>
      <Key>icons/folder.gif</Key>
      <Generation>1380149859530000</Generation>
      <MetaGeneration>2</MetaGeneration>
      <LastModified>2013-09-25T22:57:39.349Z</LastModified>
      <ETag>"c0d96102715c4916b872f91f5bf9b12c"</ETag>
      <Size>7262134</Size>
      <Owner/>
    </Contents>
  </ListBucketResult>
  `)

	in := bytes.NewReader(data)
	list, err := ParseXML(in)
	if err != nil {
		t.Fatalf("fatal")
	}

	length := list.Len()
	if length != 0 {
		t.Fatalf("length missmatch, %d", length)
	}
}
