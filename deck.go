//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Rank represents the rank of a card
type Rank int

// Suit represents the suit of a card
type Suit int

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = []Suit{Spade, Diamond, Club, Heart}

// Card represents a deck's card
type Card struct {
	Rank
	Suit
}

// String prints string representations of Ranks and Suits
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New returns a new deck of cards, also accepts functional options
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

// DefaultSort is a functional option for New, sorting the deck
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))

	return cards
}

// Sort is a functional option for New, accepting a custom sorting function
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))

		return cards
	}
}

// Less implements sorting for DefaultSort
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Shuffle is a functional option for New, shuffling a newly created deck of
// cards
func Shuffle(cards []Card) []Card {
	result := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))

	for i, j := range perm {
		result[i] = cards[j]
	}

	return result
}

// Jokers is a functional option for New, adding Jokers on a newly created deck
// of cards
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}

		return cards
	}
}

// Filter is a functional option for New, accepting a custom filtering function
// and returning a filtered deck of cards
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var result []Card

		for _, c := range cards {
			if !f(c) {
				result = append(result, c)
			}
		}

		return result
	}
}

// Deck is a functional option for New, adding additional decks on a deck of
// cards
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var result []Card

		for i := 0; i < n; i++ {
			result = append(result, cards...)
		}

		return result
	}
}
