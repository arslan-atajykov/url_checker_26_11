package checker

import (
	"net/http"
	"time"
)

func CheckURL(url string) string {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "not available"
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return "available"
	}
	return "not available"
}
