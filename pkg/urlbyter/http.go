package urlbyter

import (
	"net/http"
)

type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

