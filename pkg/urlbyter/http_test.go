package urlbyter_test

import (
	"bytes"
	"github.com/thmhoag/urlbyter/pkg/urlbyter"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MockHTTPClient struct {
	responses map[string]*http.Response
}

func NewMockHTTPClient(fakeURL string, resp []byte) urlbyter.HTTPClient {
	return &MockHTTPClient{
		responses: map[string]*http.Response{
			fakeURL: &http.Response{
				Status: strconv.Itoa(http.StatusNotFound),
				Body: ioutil.NopCloser(bytes.NewReader(resp)),
			},
		},
	}
}

func (fc *MockHTTPClient) Get(url string) (resp *http.Response, err error) {
	foundResp := fc.responses[url]
	if foundResp == nil {
		return notFound(), nil
	}

	return foundResp, nil
}

func notFound() *http.Response {
	return &http.Response{
		Status: strconv.Itoa(http.StatusNotFound),
		Body: ioutil.NopCloser(bytes.NewReader([]byte("not found"))),
	}
}