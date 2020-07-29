package urlbyter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
)

// Client provides methods for getting URL information
type Client struct {
	http HTTPClient
}

func NewClient(httpClient HTTPClient) *Client {
	return &Client{
		http: httpClient,
	}
}

// GetBytesForURL returns a human-readable string representing the byte amount
// of the body of the page located at the URL
func (c *Client) GetBytesForURL(url *url.URL) (string, error) {
	if url == nil {
		return "", errors.New("url may not be nil")
	}

	resp, err := c.http.Get(url.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return byteCount(int64(len(body))), nil
}

func byteCount(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}