package chapter09

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeUrl(raw string) string {
	if raw == "" {
		panic("url required")
	}
	parsed, err := url.Parse(strings.TrimSpace(raw))
	if err != nil || !parsed.IsAbs() {
		panic("invalid url")
	}
	scheme := strings.ToLower(parsed.Scheme)
	if scheme == "" {
		scheme = "http"
	}
	host := strings.ToLower(parsed.Hostname())
	if host == "" {
		panic("invalid url")
	}
	defaultPort := "80"
	if scheme == "https" {
		defaultPort = "443"
	}
	netloc := host
	port := parsed.Port()
	if port != "" && port != defaultPort {
		netloc = host + ":" + port
	}
	path := parsed.Path
	if path == "" {
		path = "/"
	}
	if path != "/" && strings.HasSuffix(path, "/") {
		path = strings.TrimRight(path, "/")
	}
	return scheme + "://" + netloc + path
}

type UrlDeduplicator struct {
	seen map[string]struct{}
}

func NewUrlDeduplicator() *UrlDeduplicator {
	return &UrlDeduplicator{seen: map[string]struct{}{}}
}

func (d *UrlDeduplicator) Add(raw string) bool {
	key := NormalizeUrl(raw)
	if _, ok := d.seen[key]; ok {
		return false
	}
	d.seen[key] = struct{}{}
	return true
}

func RunQ904() {
	left := NormalizeUrl("HTTP://Example.com/")
	right := NormalizeUrl("http://example.com")
	fmt.Printf("%s == %s: %t\n", left, right, left == right)
	deduplicator := NewUrlDeduplicator()
	fmt.Printf("New: %t\n", deduplicator.Add("http://Foo.com/"))
	fmt.Printf("Dup: %t\n", deduplicator.Add("http://foo.com"))
}
