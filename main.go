package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Example: albumURL := "https://www.erome.com/a/Ijac718O"
  albumURL := "https://www.erome.com/a/*****"

	htmlContent, err := fetchHTML(albumURL)
	if err != nil {
		fmt.Printf("Error fetching HTML: %v\n", err)
		return
	}

	videoURLs := extractVideoURLs(htmlContent)
	if len(videoURLs) == 0 {
		fmt.Println("No video URLs found")
		return
	}

	fmt.Printf("Found %d video(s)\n", len(videoURLs))

	for i, url := range videoURLs {
		fmt.Printf("Downloading video %d/%d: %s\n", i+1, len(videoURLs), url)
		err := downloadVideo(url, i+1)
		if err != nil {
			fmt.Printf("Error downloading video %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Video %d downloaded successfully\n", i+1)
		}
	}
}

func fetchHTML(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "ja,en;q=0.9,en-GB;q=0.8,en-US;q=0.7")
	req.Header.Set("Referer", "https://hu.erome.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func extractVideoURLs(html string) []string {
	var urls []string
	seen := make(map[string]bool)

	re := regexp.MustCompile(`<source src="(https://v\d+\.erome\.com[^"]+_720p\.mp4)"`)
	matches := re.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		if len(match) > 1 {
			url := match[1]
			if !seen[url] {
				seen[url] = true
				urls = append(urls, url)
			}
		}
	}

	return urls
}

func downloadVideo(url string, index int) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "identity;q=1, *;q=0")
	req.Header.Set("Accept-Language", "ja,en;q=0.9,en-GB;q=0.8,en-US;q=0.7")
	req.Header.Set("Range", "bytes=0-")
	req.Header.Set("Referer", "https://hu.erome.com/")
	req.Header.Set("Sec-Fetch-Dest", "video")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]
	nameWithoutExt := strings.TrimSuffix(filename, "_720p.mp4")

	finalFilename := fmt.Sprintf("video_%d_%s.mp4", index, nameWithoutExt)

	out, err := os.Create(finalFilename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Content-Length: %d\n", resp.ContentLength)
	fmt.Printf("Saved as: %s\n", finalFilename)

	return nil
}
