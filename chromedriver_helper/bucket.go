package chromedriver_helper

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
)

var BucketURL = "http://chromedriver.storage.googleapis.com/"
var NoReleases = errors.New("No releases")
var ReleaseKeyRegexp = regexp.MustCompile("\\A(\\d+)\\.(\\d+)/chromedriver_([a-z0-9]+)\\.zip\\z")
var ReleaseKeyValidationFailed = errors.New("Non release key")

type Bucket struct {
	Releases []*Release
}

type Release struct {
	Key      string
	Platform string
	Version  *ReleaseVersion
}

type ReleaseVersion struct {
	Major int
	Minor int
}

func LatestRelease(platform string) (*Release, error) {
	bucket, err := loadBucket(platform)
	if err != nil {
		return nil, err
	}

	releases := bucket.Releases
	if len(releases) == 0 {
		return nil, NoReleases
	}

	return releases[-1], nil
}

func loadBucket(platform string) (*Bucket, error) {
	bucket := &Bucket{}
	err := bucket.load(platform)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (bucket *Bucket) load(platform string) error {
	contents, err := loadContents()
	if err != nil {
		return nil, err
	}

	for _, c := range contents {
		release, err := loadReleaseFromKey(c.Key)
		if err != nil {
			if err == ReleaseKeyValidationFailed {
				continue
			} else {
				return err
			}
		}
		if release.Platform != platform {
			continue
		}
		bucket.Releases = append(bucket.Releases, release)
	}

	sort.Sort(bucket)
	return nil
}

func loadReleaseFromKey(key string) (*Release, error) {
	matches := ReleaseKeyRegexp.FindStringSubmatch(key)
	if matches == nil {
		return nil, ReleaseKeyValidatinoFailed
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	platform := matches[3]

	return &Release{
		Key:      key,
		Platform: platform,
		Version: &ReleaseVersion{
			Major: major,
			Minor: minor,
		},
	}, nil
}

func (release *Release) URL() string {
	return BucketURL + release.Key
}
