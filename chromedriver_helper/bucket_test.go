package chromedriver_helper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createServer(body string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, body)
	})

	return httptest.NewServer(handler)
}

func TestBucketSingleContents(t *testing.T) {
	s := createServer(`
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
	defer s.Close()

	b := NewBucketWithURL(s.URL)
	release, err := b.LatestRelease("linux32")
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, release)

	assert.Equal(t, 2, release.MajorVersion)
	assert.Equal(t, 0, release.MinorVersion)
	assert.Equal(t, "2.0/chromedriver_linux32.zip", release.Key)
}

func TestBucketMultiPlatformContents(t *testing.T) {
	s := createServer(`
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
				<Contents>
					<Key>2.0/chromedriver_linux64.zip</Key>
					<Generation>1380149859530000</Generation>
					<MetaGeneration>2</MetaGeneration>
					<LastModified>2013-09-25T22:57:39.349Z</LastModified>
					<ETag>"c0d96102715c4916b872f91f5bf9b12c"</ETag>
					<Size>7262134</Size>
					<Owner/>
				</Contents>
			</ListBucketResult>
		`)
	defer s.Close()

	b := NewBucketWithURL(s.URL)
	release, err := b.LatestRelease("linux32")
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, release)

	assert.Equal(t, 2, release.MajorVersion)
	assert.Equal(t, 0, release.MinorVersion)
	assert.Equal(t, "2.0/chromedriver_linux32.zip", release.Key)
}

func TestBucketMultiVersionContents(t *testing.T) {
	s := createServer(`
			<?xml version='1.0' encoding='UTF-8'?>
			<ListBucketResult xmlns="http://doc.s3.amazonaws.com/2006-03-01">
				<Name>chromedriver</Name>
				<Prefix/>
				<Marker/>
				<IsTruncated>false</IsTruncated>
				<Contents>
					<Key>2.1/chromedriver_linux32.zip</Key>
					<Generation>1380149859530000</Generation>
					<MetaGeneration>2</MetaGeneration>
					<LastModified>2013-09-25T22:57:39.349Z</LastModified>
					<ETag>"c0d96102715c4916b872f91f5bf9b12c"</ETag>
					<Size>7262134</Size>
					<Owner/>
				</Contents>
				<Contents>
					<Key>2.10/chromedriver_linux32.zip</Key>
					<Generation>1380149859530000</Generation>
					<MetaGeneration>2</MetaGeneration>
					<LastModified>2013-09-25T22:57:39.349Z</LastModified>
					<ETag>"c0d96102715c4916b872f91f5bf9b12c"</ETag>
					<Size>7262134</Size>
					<Owner/>
				</Contents>
			</ListBucketResult>
		`)
	defer s.Close()

	b := NewBucketWithURL(s.URL)
	release, err := b.LatestRelease("linux32")
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.NotNil(t, release)

	assert.Equal(t, 2, release.MajorVersion)
	assert.Equal(t, 10, release.MinorVersion)
	assert.Equal(t, "2.10/chromedriver_linux32.zip", release.Key)
}
