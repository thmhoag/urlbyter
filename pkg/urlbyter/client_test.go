package urlbyter_test

import (
	"fmt"
	"github.com/thmhoag/urlbyter/pkg/urlbyter"
	"io/ioutil"
	"net/url"
	"testing"
)

func TestClient_GetBytesForURL(t *testing.T) {
	type fields struct {
		http urlbyter.HTTPClient
	}
	type args struct {
		url *url.URL
	}
	type test struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}
	var tests []test

	files := []struct {
		path string
		url string
		expected string
	}{
		{
			path: "./test/github_explore_test.html",
			url: "https://github.com/explore",
			expected: "418.5 kB",
		},
		{
			path: "./test/github_test.html",
			url: "https://github.com",
			expected: "136.4 kB",
		},
		{
			path: "./test/google_search_golang_test.html",
			url: "https://www.google.com/search?q=golang",
			expected: "6.3 kB",
		},
		{
			path: "./test/google_search_k8s_test.html",
			url: "https://www.google.com/search?q=k8s",
			expected: "6.3 kB",
		},
		{
			path: "./test/google_search_salesforce_test.html",
			url: "https://www.google.com/search?q=salesforce",
			expected: "6.3 kB",
		},
		{
			path: "./test/google_test.html",
			url: "https://www.google.com",
			expected: "12.0 kB",
		},
		{
			path: "./test/microsoft_365_test.html",
			url: "https://www.microsoft.com/en-us/microsoft-365",
			expected: "177.2 kB",
		},
		{
			path: "./test/microsoft_enus_test.html",
			url: "https://www.microsoft.com/en-us/",
			expected: "199.2 kB",
		},
		{
			path: "./test/microsoft_test.html",
			url: "https://www.microsoft.com/",
			expected: "1.0 kB",
		},
	}

	for _, f := range files {
		bytes, err := ioutil.ReadFile(f.path)
		if err != nil {
			t.Fatalf("unable to read test file: %s\n", f.path)
		}

		url, err := url.Parse(f.url)
		if err != nil {
			t.Fatalf("invalid URL test parameter: %s", f.url)
		}

		tests = append(tests, test{
			name: fmt.Sprintf("when test url is %s", f.url),
			fields: fields{
				http: NewMockHTTPClient(f.url, bytes),
			},
			args: args{
				url: url,
			},
			want: f.expected,
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := urlbyter.NewClient(tt.fields.http)
			got, err := c.GetBytesForURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBytesForURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBytesForURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}