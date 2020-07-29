package urlbyter

import (
	"bufio"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type BatchOpts struct {
	Client *Client
}

type BatchRunner struct {
	opts *BatchOpts
}

type BatchResult struct {
	Host string
	URL  *url.URL
	Bytes string
	Err error
}

func NewBatchRunner(opts *BatchOpts) *BatchRunner {
	if opts == nil {
		panic("invalid BatchOpts")
	}

	if opts.Client == nil {
		opts.Client = NewClient(http.DefaultClient)
	}

	return &BatchRunner{
		opts: opts,
	}
}

func (br *BatchRunner) ProcessFile(path string) ([]*BatchResult, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rawFileURLs []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		rawFileURLs = append(rawFileURLs, s.Text())
	}

	// parsedURLs won't return any invalid ones
	parsedURLs := ParseURLs(rawFileURLs)
	if len(parsedURLs) < 1 {
		return []*BatchResult{}, nil
	}

	// threadsafe destination for results from URLs
	resultChan := make(chan *BatchResult, len(parsedURLs))

	// make sure we can wait out all the goroutines
	var wg sync.WaitGroup

	// group URLs by hostname
	groupedURLs := GroupURLs(parsedURLs)

	// process URLs by hostname "group" so we don't
	// perform more than 1 request to the same host at
	// at a time
	for _, urls := range groupedURLs {

		// add 1 to the WaitGroup for this goroutine
		wg.Add(1)

		// execute the goroutine
		go func(urls []*url.URL) {
			defer wg.Done()

			for i := range urls {
				u := urls[i]
				res := &BatchResult{
					Host: u.Hostname(),
					URL: u,
				}

				resp, err := br.opts.Client.GetBytesForURL(u)
				if err != nil {
					res.Err = err
					continue
				}

				res.Bytes = resp
				resultChan <- res
			}
		}(urls)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// close result channel to guarantee we can read all of it
	// we know we're done at this point anyway
	close(resultChan)

	// cleanup results into a slice we can return
	var results []*BatchResult
	for result := range resultChan {
		results = append(results, result)
	}

	return results, nil
}