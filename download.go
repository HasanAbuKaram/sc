package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

func downloadFile(filepath string, url string) error {
	// Check if file already exists
	if _, err := os.Stat(filepath); err == nil {
		return fmt.Errorf("file %s already exists", filepath)
	}

	// Check if directory exists, if not create it
	dir := path.Dir(filepath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create a custom http client
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 60 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	// Get the data
	resp, err := netClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
