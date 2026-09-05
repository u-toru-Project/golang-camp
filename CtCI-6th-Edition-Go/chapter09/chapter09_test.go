package chapter09

import (
	"strings"
	"testing"
)

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	fn()
}

func TestStockSocialCrawlerUrls(t *testing.T) {
	aggregator := NewStockAggregator()
	aggregator.OnTick(Tick{"AAPL", 100.0, 10})
	aggregator.OnTick(Tick{"AAPL", 105.0, 5})
	aggregator.OnTick(Tick{"AAPL", 98.0, 2})
	bar := aggregator.GetOhlc("AAPL")
	if bar == nil || bar.Open != 100 || bar.High != 105 || bar.Low != 98 || bar.Close != 98 || bar.Volume != 17 {
		t.Fatal(bar)
	}
	if NewStockAggregator().GetOhlc("MSFT") != nil {
		t.Fatal("missing")
	}
	mustPanic(t, func() { NewStockAggregator().OnTick(Tick{"AAPL", 0, 1}) })
	mustPanic(t, func() { NewStockAggregator().OnTick(Tick{"AAPL", 10, -1}) })

	network := NewSocialNetwork()
	for _, user := range []string{"a", "b", "c", "d"} {
		network.AddUser(user)
	}
	network.Follow("a", "b")
	network.Follow("b", "c")
	network.Follow("a", "d")
	fof := network.FriendsOfFriends("a")
	found := false
	for _, u := range fof {
		if u == "c" {
			found = true
		}
	}
	if !found {
		t.Fatal(fof)
	}
	mustPanic(t, func() { NewSocialNetwork().Follow("a", "a") })
	mustPanic(t, func() { NewSocialNetwork().FriendsOfFriends("missing") })

	crawler := NewWebCrawler(MockGraph)
	if strings.Join(crawler.Crawl("http://a"), ",") != "http://a,http://b,http://c,http://d" {
		t.Fatal(crawler.Crawl("http://a"))
	}
	limited := NewWebCrawler(MockGraph, 2)
	order := limited.Crawl("http://a")
	if len(order) != 2 || order[0] != "http://a" || order[1] != "http://b" {
		t.Fatal(order)
	}
	mustPanic(t, func() { NewWebCrawler(MockGraph).Crawl("") })

	if NormalizeUrl("HTTP://Example.com/") != NormalizeUrl("http://example.com") {
		t.Fatal("normalize host")
	}
	if NormalizeUrl("https://example.com:443/path/") != NormalizeUrl("https://example.com/path") {
		t.Fatal("normalize path")
	}
	deduplicator := NewUrlDeduplicator()
	if !deduplicator.Add("http://Foo.com/") || deduplicator.Add("http://foo.com") {
		t.Fatal("dedup")
	}
}

func TestCacheSalesFinancePaste(t *testing.T) {
	cache := NewLruCache[string, int](2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	a, foundA := cache.TryGet("a")
	if !foundA || a != 1 {
		t.Fatal(foundA, a)
	}
	cache.Put("c", 3)
	if _, hasB := cache.TryGet("b"); hasB {
		t.Fatal("b should be evicted")
	}
	a2, foundA2 := cache.TryGet("a")
	if !foundA2 || a2 != 1 || cache.Count() != 2 {
		t.Fatal(foundA2, a2, cache.Count())
	}
	mustPanic(t, func() { NewLruCache[string, int](0) })

	rank := NewSalesRank(2)
	rank.RecordSale("p1", 5)
	rank.RecordSale("p2", 10)
	rank.RecordSale("p3", 7)
	top := rank.TopK()
	if len(top) != 2 || top[0].ProductId != "p2" || top[0].Count != 10 || top[1].ProductId != "p3" || top[1].Count != 7 {
		t.Fatal(top)
	}
	mustPanic(t, func() { NewSalesRank(1).RecordSale("p1", 0) })

	manager := NewPersonalFinancialManager()
	manager.AddRule("grocery", "food")
	manager.AddRule("uber", "transport")
	transactions := []Transaction{{-50.0, "Whole Foods grocery"}, {-20.0, "Uber ride"}}
	if manager.Categorize(transactions[0]) != "food" {
		t.Fatal("cat")
	}
	summary := manager.Summary(transactions)
	if summary["food"] != -50 || summary["transport"] != -20 {
		t.Fatal(summary)
	}
	if NewPersonalFinancialManager().Categorize(Transaction{5, "coffee"}) != "uncategorized" {
		t.Fatal("uncat")
	}

	pasteBin := NewPasteBin(60)
	id := pasteBin.Paste("hello world")
	got := pasteBin.Get(id)
	if got == nil || *got != "hello world" {
		t.Fatal(got)
	}
	clock := 0.0
	timed := NewPasteBin(1.0, func() float64 { return clock })
	id2 := timed.Paste("x")
	if g := timed.Get(id2); g == nil || *g != "x" {
		t.Fatal(g)
	}
	clock = 2.0
	if timed.Get(id2) != nil {
		t.Fatal("expired")
	}
	clock = 0.0
	p2 := NewPasteBin(1.0, func() float64 { return clock })
	p2.Paste("keep")
	clock = 2.0
	if p2.PurgeExpired() != 1 || p2.PurgeExpired() != 0 {
		t.Fatal("purge")
	}
}
