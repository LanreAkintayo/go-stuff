package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
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
		os.Remove(filePath)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		os.Remove(filePath)
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

type Result struct {
	URL      string
	Filename string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, destDir string, maxConcurrent int) error {
	// Make sure dest dir exist
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		return err
	}
	rateLimiter := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup

	results := make(chan Result)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			rateLimiter <- struct{}{}
			defer func() { <-rateLimiter }()

			start := time.Now()
			// Then download the file;
			fileName := filepath.Base(url)
			filePath := filepath.Join(destDir, fileName)

			out, err := os.Create(filePath)
			if err != nil {
				// fmt.Println("Error creating file: ", err);
				results <- Result{
					URL:   url,
					Error: err,
				}
				return
			}
			defer out.Close()

			resp, err := http.Get(url)
			if err != nil {
				os.Remove(filePath)
				// fmt.Println("Error downloading: ", err);
				results <- Result{
					URL:   url,
					Error: err,
				}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				os.Remove(filePath)
				// fmt.Println("Bad status: ", resp.Status);
				results <- Result{
					URL:   url,
					Error: fmt.Errorf("Bad status: %s", resp.Status),
				}
				return
			}

			size, err := io.Copy(out, resp.Body)
			if err != nil {
				os.Remove(filePath)
				// fmt.Println("Error copying: ", err);
				results <- Result{
					URL:   url,
					Error: err,
				}
				return
			}

			duration := time.Since(start)

			results <- Result{
				URL:      url,
				Filename: fileName,
				Size:     size,
				Duration: duration,
				Error:    nil,
			}
			// fmt.Println("Downloaded in: ", duration.Seconds(), "seconds")

		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error

	start := time.Now()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error downloading %s : %s\n", result.URL, result.Error.Error())
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("Downloded %s (%d bytes) in %s\n", result.Filename, result.Size, result.Duration.Seconds())
		}
	}

	startedSince := time.Since(start)

	fmt.Printf("All downloads completed in %s, Total %d bytes\n", startedSince.String(), totalSize)
	if len(errors) > 0 {
		fmt.Printf("Encountered %d errors\n", len(errors))
	}

	return nil
}

func main() {
	// url := "https://go.dev/images/go_amex_case_study.png"
	urls := []string{
		"https://go.dev/images/favicon-gopher.png",
		"https://go.dev/images/go-logo-blue.svg",
	}

	// err := DownloadFile(url, "./file_downloader/")
	// err := SequentialDownloader(urls, "./file_downloader/")
	err := ConcurrentDownloader(urls, "./file_downloader/concurrent/", 3)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Success!")
}
