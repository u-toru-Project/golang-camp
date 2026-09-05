package chapter09

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

type pasteEntry struct {
	content string
	created float64
}

type PasteBin struct {
	ttlSeconds float64
	now        func() float64
	store      map[string]pasteEntry
}

func NewPasteBin(ttlSec float64, now ...func() float64) *PasteBin {
	if ttlSec <= 0 {
		panic("ttl must be positive")
	}
	clock := func() float64 { return float64(time.Now().UnixNano()) / 1e9 }
	if len(now) > 0 && now[0] != nil {
		clock = now[0]
	}
	return &PasteBin{ttlSeconds: ttlSec, now: clock, store: map[string]pasteEntry{}}
}

func (p *PasteBin) Paste(content string) string {
	if content == "" {
		panic("content required")
	}
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	pasteId := strings.TrimRight(strings.NewReplacer("+", "-", "/", "_").Replace(base64.StdEncoding.EncodeToString(buf)), "=")
	p.store[pasteId] = pasteEntry{content, p.now()}
	return pasteId
}

func (p *PasteBin) Get(pasteId string) *string {
	entry, ok := p.store[pasteId]
	if !ok {
		return nil
	}
	if p.now()-entry.created > p.ttlSeconds {
		delete(p.store, pasteId)
		return nil
	}
	content := entry.content
	return &content
}

func (p *PasteBin) PurgeExpired() int {
	now := p.now()
	expired := make([]string, 0)
	for id, entry := range p.store {
		if now-entry.created > p.ttlSeconds {
			expired = append(expired, id)
		}
	}
	for _, id := range expired {
		delete(p.store, id)
	}
	return len(expired)
}

func RunQ908() {
	pasteBin := NewPasteBin(60.0)
	id := pasteBin.Paste("hello world")
	fmt.Printf("%s: %s\n", id, *pasteBin.Get(id))
}
