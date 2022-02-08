package tool

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var c http.Client

func init() {
	jar, _ := cookiejar.New(nil)
	c = http.Client{
		Timeout: time.Duration(60) * time.Second,
		Jar:     jar,
	}
}

func Get(url string) (io.ReadCloser, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
	req.Header.Set("referer", "https://jable.tv/")
	req.Header.Set("origin", "https://jable.tv/")

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
