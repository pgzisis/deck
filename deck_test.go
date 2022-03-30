package deck

import (
	"math/rand"
	"testing"
)

// TestNew should create a new deck
func TestNew(t *testing.T) {
	deck := New()
	want := 52
	got := len(deck)

	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}

// TestNewDefaultSort should create a new deck with default sorting
func TestNewDefaultSort(t *testing.T) {
	got := New(DefaultSort)[0]
	want := Card{Suit: Spade, Rank: Ace}

	if got != want {
		t.Errorf("Want %v got %v", want, got)
	}
}

// TestSort should create a new deck with custom sorting
func TestSort(t *testing.T) {
	got := New(Sort(Less))[0]
	want := Card{Suit: Spade, Rank: Ace}

	if got != want {
		t.Errorf("Want %v got %v", want, got)
	}
}

// TestShuffle should create a shuffled deck
func TestShuffle(t *testing.T) {
	// Make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40 35 ...]
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]
	d := New(Shuffle)

	if d[0] != first {
		t.Errorf("Want %v got %v", first, d[0])
	}

	if d[1] != second {
		t.Errorf("Want %v got %v", second, d[1])
	}
}

// TestJokers should add a specified number of Jokers to a newly created deck
func TestJokers(t *testing.T) {
	want := 3
	d := New(Jokers(want))
	got := 0
	for _, c := range d {
		if c.Suit == Joker {
			got++
		}
	}

	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}

// TestFilter should create a new deck with custom filtering
func TestFilter(t *testing.T) {
	f := func(card Card) bool {
		return card.Rank < 5
	}
	d := New(Filter(f))
	want := 52 - 4*4
	got := len(d)

	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}

// TestDeck should create a deck with a user specified number of decks
func TestDeck(t *testing.T) {
	suits := 4
	ranks := 13
	decks := 3
	d := New(Deck(decks))
	want := suits * ranks * decks
	got := len(d)

	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}
