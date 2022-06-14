// Shuffle deck of cards
package main

import (
	"example.com/nodequeue"
	"math/rand"
	"time"
	"fmt"
)

type Card struct {
	Rank string
	Suit string
}

type Deck struct {
	Cards []Card
}

var ranks = []string {"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var suits = []rune {'\u2660', '\u2661', '\u2662', '\u2663'}

func NewDeck() (deck Deck) {
	for _, suit := range(suits) {
		for _, rank := range(ranks) {
			deck.Cards = append(deck.Cards, Card{rank, string(suit)})
		}
	}
	return deck
}

func (deck Deck) Shuffle() Deck {
	q1 := nodequeue.Queue[Card]{}
	q2 := nodequeue.Queue[Card]{}
	// Cut deck 
	mismatch := -5 + rand.Intn(11) // -5 to 5
	var i int
	for i = 0; i < 26 + mismatch; i++ {
		q1.Insert(deck.Cards[i])
	}
	for ; i < 52; i++ {
		q2.Insert(deck.Cards[i])
	}
	// Rebuild deck
	deck = Deck{}
	for {
		if q1.Size() == 0 || q2.Size() == 0 {
			break
		}
		card := q1.Remove()
		deck.Cards = append(deck.Cards, card)
		card = q2.Remove()
		deck.Cards = append(deck.Cards, card)
	}
	if q2.Size() == 0 {
		for {
			if q1.Size() == 0 {
				break
			}
			card := q1.Remove()
			deck.Cards = append(deck.Cards, card)
		}
	}
	if q1.Size() == 0 {
		for {
			if q2.Size() == 0 {
				break
			}
			card := q2.Remove()
			deck.Cards = append(deck.Cards, card)
		}
	}
	return deck
}

func main() {
	rand.Seed(time.Now().UnixNano())
	deck := NewDeck()
	fmt.Println("\nOriginal deck: ", deck)
	// Cut deck 5 times
	for index := 0; index < 5; index++ {
		deck = deck.Shuffle()
	}
	fmt.Println("\nShuffled deck: ", deck)
}


