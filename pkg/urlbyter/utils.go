package urlbyter

import (
	"fmt"
	"net/url"
)

// ParseURLs parses strings to the url.URL struct ignoring
// any that are not valid urls
func ParseURLs(urls []string) []*url.URL {
	var results []*url.URL
	for _, u := range urls {
		newURL, _ := url.Parse(u)
		results = append(results, newURL)
	}

	return results
}

// GroupURLs groups a collection of URLs by their hostname
// Will panic on invalid URLs
func GroupURLs(urls []*url.URL) map[string][]*url.URL {
	results := make(map[string][]*url.URL)
	for i := range urls {
		u := urls[i]

		// Making a conscious decision here to group urls with different ports under
		// the same hostname
		hostname := u.Hostname()
		if hostname == "" {
			// panic because these URLs should have been validated already
			panic(fmt.Sprintf("invalid URL %s", u.String()))
		}

		results[hostname] = append(results[hostname], u)
	}

	return results
}