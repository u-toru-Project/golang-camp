package chapter09

import "fmt"

type WebCrawler struct {
	fetch    func(string) []string
	maxPages int
}

func NewWebCrawler(fetch func(string) []string, maxPages ...int) *WebCrawler {
	m := 100
	if len(maxPages) > 0 {
		m = maxPages[0]
	}
	return &WebCrawler{fetch: fetch, maxPages: m}
}

func (c *WebCrawler) Crawl(startUrl string) []string {
	if startUrl == "" {
		panic("start_url required")
	}
	visited := map[string]struct{}{}
	order := make([]string, 0)
	queue := []string{startUrl}
	for len(queue) > 0 && len(order) < c.maxPages {
		url := queue[0]
		queue = queue[1:]
		if _, ok := visited[url]; ok {
			continue
		}
		visited[url] = struct{}{}
		order = append(order, url)
		for _, link := range c.fetch(url) {
			if _, ok := visited[link]; !ok {
				queue = append(queue, link)
			}
		}
	}
	return order
}

func MockGraph(url string) []string {
	graph := map[string][]string{
		"http://a": {"http://b", "http://c"},
		"http://b": {"http://d"},
		"http://c": {},
		"http://d": {},
	}
	if links, ok := graph[url]; ok {
		return links
	}
	return nil
}

func RunQ903() {
	crawler := NewWebCrawler(MockGraph)
	fmt.Println(crawler.Crawl("http://a"))
}
