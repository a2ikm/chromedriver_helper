package chromedriver_helper

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	DefaultBucket = NewBucket("http://chromedriver.storage.googleapis.com/")
)

var (
	NoReleases = errors.New("No Releases")
)

type Bucket struct {
	URL      string
	Releases Releases
}

type Releases []*Release

type Release struct {
	Bucket       *Bucket
	MajorVersion int
	MinorVersion int
	Platform     string
}

func NewBucket(url string) *Bucket {
	return &Bucket{
		URL: url,
	}
}

func LatestRelease(platform string) (*Release, error) {
	return DefaultBucket.LatestRelease(platform)
}

func (b *Bucket) LatestRelease(platform string) (*Release, error) {
	releases, err := b.loadReleases(platform)
	if err != nil {
		return nil, err
	}

	length := len(releases)
	if length == 0 {
		return nil, NoReleases
	}

	return releases[length-1], nil
}

type listBucketResult struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string
	Prefix      string
	Marker      string
	IsTruncated string
	Contents    []contents
}

type contents struct {
	XMLName        xml.Name `xml:"Contents"`
	Key            string
	Generation     int
	MetaGeneration int
	LastModified   string
	ETag           string
	Size           int
}

func (b *Bucket) loadReleases(platform string) (Releases, error) {
	res, err := http.Get(b.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := listBucketResult{}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	keyRegexpStr := "\\A(\\d+)\\.(\\d+)/chromedriver_" + platform + "\\.zip\\z"
	keyRegexp := regexp.MustCompile(keyRegexpStr)

	var releases Releases
	for _, c := range result.Contents {
		matches := keyRegexp.FindStringSubmatch(c.Key)
		if matches == nil {
			continue
		}

		major, _ := strconv.Atoi(matches[1])
		minor, _ := strconv.Atoi(matches[2])

		release := &Release{
			Bucket:       b,
			MajorVersion: major,
			MinorVersion: minor,
			Platform:     platform,
		}
		releases = append(releases, release)
	}

	sort.Sort(releases)
	b.Releases = releases
	return releases, nil
}

func (r *Release) URL() string {
	return strings.Join([]string{r.Bucket.URL, r.Key()}, "")
}

func (r *Release) Key() string {
	return fmt.Sprintf("%d.%d/chromedriver_%s.zip", r.MajorVersion, r.MinorVersion, r.Platform)
}

func (s Releases) Len() int {
	return len(s)
}

func (s Releases) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Releases) Less(i, j int) bool {
	ri, rj := s[i], s[j]
	if ri.MajorVersion == rj.MajorVersion {
		return ri.MinorVersion < rj.MinorVersion
	} else {
		return ri.MajorVersion < rj.MajorVersion
	}
}
