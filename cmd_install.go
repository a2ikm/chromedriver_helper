package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/a2ikm/chromedriver_helper/chromedriver_helper"
)

type CmdInstall struct{}

func (c *CmdInstall) Help() string {
	return "Install latest binary"
}

func (c *CmdInstall) Run(args []string) int {
	err := c.realRun()
	if err != nil {
		log.Fatalln(err.Error())
		return 1
	}

	return 0
}

func (c *CmdInstall) Synopsis() string {
	return ""
}

func (c *CmdInstall) realRun() error {
	fmt.Println("Getting released versions")

	release, err := c.targetRelease()
	if err != nil {
		return err
	}

	fmt.Printf("Installing %s\n", release.Key())

	rc, err := c.download(release)
	if err != nil {
		return err
	}
	defer rc.Close()

	return c.install(rc)
}

func (c *CmdInstall) targetRelease() (*chromedriver_helper.Release, error) {

	platform, err := chromedriver_helper.Platform()
	if err != nil {
		return nil, err
	}

	release, err := chromedriver_helper.LatestRelease(platform)
	if err != nil {
		return nil, err
	}

	return release, nil
}

func (c *CmdInstall) download(release *chromedriver_helper.Release) (io.ReadCloser, error) {
	url := release.URL()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	readerAt := bytes.NewReader(b)
	reader, err := zip.NewReader(readerAt, res.ContentLength)
	if err != nil {
		return nil, err
	}

	f := reader.File[0]
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}

	return rc, nil
}

func (c *CmdInstall) install(rc io.ReadCloser) error {
	err := chromedriver_helper.MakeInstallDir()
	if err != nil {
		return err
	}

	path, err := chromedriver_helper.BinaryPath()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, rc)
	if err != nil {
		return err
	}

	return nil
}
