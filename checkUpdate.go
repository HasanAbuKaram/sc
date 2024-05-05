package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		BrowserDownloadUrl string `json:"browser_download_url"`
	} `json:"assets"`
}

func checkForUpdates() {
	resp, err := http.Get(repoUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if release.TagName != version {
		// Check if Assets is not empty
		if len(release.Assets) == 0 {
			fmt.Println("No assets found in the release.")
			return
		}

		var url = release.Assets[0].BrowserDownloadUrl

		// Get the system's temporary directory
		tempDir := os.TempDir()

		// Create a hidden directory in the temp directory
		hiddenDir := filepath.Join(tempDir, ".hidden")
		os.MkdirAll(hiddenDir, os.ModePerm)

		// Download the file to the hidden directory
		err := downloadFile(filepath.Join(hiddenDir, fmt.Sprintf("%v.exe", release.TagName)), url)
		if err != nil {
			fmt.Println("Error downloading file:", err)
		} else {
			fmt.Println("File downloaded and installed successfully.")
			// Install the downloaded binary
			if err := installBinary(filepath.Join(hiddenDir, fmt.Sprintf("%v.exe", release.TagName))); err != nil {
				fmt.Println("Error installing binary:", err)
			}
		}

	} else {
		currentTime := time.Now()
		formattefTime := currentTime.Format("Monday, January 02, 2006 15:04:05")
		_, weekNumber := currentTime.ISOWeek()
		fmt.Printf("As of %s (week %02d), you are running the latest version: %s\n", formattefTime, weekNumber, version)
	}
}
