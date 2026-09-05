package chapter07

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

var Suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}

type Card struct {
	Suit string
	Rank int
}

func NewCard(suit string, rank int) Card {
	if !slices.Contains(Suits, suit) {
		panic("invalid suit")
	}
	if rank < 2 || rank > 14 {
		panic("invalid rank")
	}
	return Card{Suit: suit, Rank: rank}
}

func (c Card) String() string {
	label := fmt.Sprintf("%d", c.Rank)
	switch c.Rank {
	case 11:
		label = "J"
	case 12:
		label = "Q"
	case 13:
		label = "K"
	case 14:
		label = "A"
	}
	return label + " of " + c.Suit
}

type Deck struct {
	cards []Card
	next  func(n int) int
}

func NewDeck() *Deck {
	return NewDeckWithNext(nil)
}

func NewDeckWithNext(next func(n int) int) *Deck {
	if next == nil {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		next = func(n int) int { return r.Intn(n) }
	}
	d := &Deck{next: next}
	for _, suit := range Suits {
		for rank := 2; rank <= 14; rank++ {
			d.cards = append(d.cards, NewCard(suit, rank))
		}
	}
	return d
}

func (d *Deck) Shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := d.next(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) Deal(count int) []Card {
	if count < 0 {
		panic("count must be non-negative")
	}
	if count > len(d.cards) {
		panic("not enough cards")
	}
	dealt := append([]Card{}, d.cards[:count]...)
	d.cards = d.cards[count:]
	return dealt
}

func (d *Deck) Remaining() int { return len(d.cards) }

func RunQ701() {
	deck := NewDeckWithNext(rand.New(rand.NewSource(0)).Intn)
	deck.Shuffle()
	hand := deck.Deal(5)
	fmt.Printf("Dealt %d cards, %d remaining\n", len(hand), deck.Remaining())
	for _, card := range hand {
		fmt.Println(card)
	}
}
