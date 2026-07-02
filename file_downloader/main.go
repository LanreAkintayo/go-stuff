package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadFile(url, destDir string) error {
	fileName := filepath.Base(url)
	filePath := filepath.Join(destDir, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Downloading...")
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	duration := time.Since(start)
	fmt.Println("Downloaded in:", duration.Seconds(), "seconds")
	return nil
}

func SequentialDownloader(urls []string, destDir string) error {
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		return err
	}

	start := time.Now()

	for _, url := range urls {
		err := DownloadFile(url, destDir)
		if err != nil {
			fmt.Println("Error downloading file: ", err, url)
			continue
		}
	}

	duration := time.Since(start)
	fmt.Println("Sequential download complete in:", duration.Seconds(), "seconds")
	return nil
}

func main() {
	// url := "https://go.dev/images/go_amex_case_study.png"
	urls := []string{
		"https://go.dev/images/favicon-gopher.png",
		"https://go.dev/images/go-logo-blue.svg",
	}
	
	// err := DownloadFile(url, "./file_downloader/")
	err := SequentialDownloader(urls, "./file_downloader/")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Success!")
}
