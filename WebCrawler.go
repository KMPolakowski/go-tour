package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type safeCrawlHistory struct {
	crawledUrls map[string]bool
	mux         sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (safeCrawlHistory *safeCrawlHistory) Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	safeCrawlHistory.mux.Lock()
	_, exists := safeCrawlHistory.crawledUrls[url]

	if exists {
		safeCrawlHistory.mux.Unlock()
		return
	}
	safeCrawlHistory.crawledUrls[url] = true
	safeCrawlHistory.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- fmt.Sprintf("found: %s %q\n", url, body)

	for _, u := range urls {
		go safeCrawlHistory.Crawl(u, depth-1, fetcher, ch)
	}

	return
}

func main() {
	ch := make(chan string)
	crawls := safeCrawlHistory{crawledUrls: make(map[string]bool)}
	go crawls.Crawl("https://golang.org/", 4, fetcher, ch)

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
