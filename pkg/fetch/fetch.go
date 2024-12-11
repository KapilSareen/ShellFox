package fetch

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func Fetch(url string) (string) {
	time.Sleep(2 * time.Second)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		resp, err := http.Get("https://www.google.com/search?q="+url)
		if err != nil {
			return ""
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ""
		}
		condensedBody := strings.Join(strings.Fields(string(body)), " ")
		return condensedBody
	}
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	condensedBody := strings.Join(strings.Fields(string(body)), " ")
	return condensedBody
}
