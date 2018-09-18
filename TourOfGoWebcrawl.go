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

type synchronizedMap struct {
	urls map[string]bool
	mtx sync.Mutex
}

var syncUrls = &synchronizedMap{urls: make(map[string]bool)}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	//var ch = make(chan string)
	
	var wcrawl func(url string, depth int, fetcher Fetcher, c chan string)
	wcrawl = func(url string, depth int, fetcher Fetcher, c chan string) {
		if c != nil {
			defer close(c)
		}
		if depth <= 0 || !syncUrls.AddUrl(url) {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		
		ccg := make([]chan string, len(urls))
		for i, u := range urls {
			ccg[i]  = make(chan string)
			go wcrawl(u, depth-1, fetcher, ccg[i])
		}
		
		for i := range ccg {
			for resp := range ccg[i] {
				c <- resp
			}
		}
	}
	wcrawl(url, depth, fetcher, nil)
	
	return
}

func (smap *synchronizedMap) AddUrl(url string) bool {
	defer smap.mtx.Unlock()
	
	smap.mtx.Lock()
	
	if !smap.urls[url] {
		smap.urls[url] = true
		return true
	}
	
	return false
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
