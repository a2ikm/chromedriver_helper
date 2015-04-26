package main

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/a2ikm/chromedriver_helper/bucket"
	"github.com/a2ikm/chromedriver_helper/chromedriver_helper"
)

type CmdInstall struct{}

func (c *CmdInstall) Help() string {
	return "Install latest binary"
}

func (c *CmdInstall) Run(args []string) int {
	err := c.RealRun()
	if err != nil {
		return 1
	}

	return 0
}

func (c *CmdInstall) Synopsis() string {
	return ""
}

func (c *CmdInstall) RealRun() error {
	release, err := c.TargetRelease()
	if err != nil {
		return err
	}

	return c.Install(release)
}

func (c *CmdInstall) TargetRelease() (*bucket.Release, error) {
	platform, err := chromedriver_helper.Platform()
	if err != nil {
		return nil, err
	}

	release, err := bucket.LatestReleaseForPlatform(platform)
	if err != nil {
		return nil, err
	}

	return release, nil
}

func (c *CmdInstall) Install(release *bucket.Release) error {
	url := release.URL()
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	readerAt := bytes.NewReader(b)
	reader, err := zip.NewReader(readerAt, res.ContentLength)
	if err != nil {
		return nil
	}

	f := reader.File[0]
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	err = chromedriver_helper.PrepareDir()
	if err != nil {
		return err
	}

	path, err := chromedriver_helper.Path()
	if err != nil {
		return err
	}

	f2, err := os.OpenFile(path,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		f.Mode())
	if err != nil {
		return err
	}
	defer f2.Close()

	_, err = io.Copy(f2, rc)
	if err != nil {
		return err
	}

	return nil
}
