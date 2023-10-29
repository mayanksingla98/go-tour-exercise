package main

import (
	"fmt"
	"sync"
)

type counter struct {
	v map[string]bool
	m sync.Mutex
}

var cn = counter{v: make(map[string]bool)}
var wg = sync.WaitGroup{}

func Crawl(url string, depth int, fetcher fakeFetcher, ch chan string) {

	defer wg.Done()
	// cn.m.Lock()
	if depth <= 0 || cn.v[url] {
		return
	}
	cn.v[url] = true
	// cn.m.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, ch)
	}

	ch <- fmt.Sprintf("found: %s %q", url, body)
	return
}

func main() {
	ch := make(chan string)

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for vk := range ch {
		fmt.Println(vk)
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
