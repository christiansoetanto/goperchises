package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(
		Card{
			Suit: Heart,
			Rank: Ace,
		},
	)
	fmt.Println(
		Card{
			Suit: Diamond,
			Rank: Two,
		},
	)
	fmt.Println(

		Card{
			Suit: Club,
			Rank: King,
		},
	)
	fmt.Println(

		Card{
			Suit: Spade,
			Rank: Queen,
		},
	)
	fmt.Println(

		Card{
			Suit: Joker,
		},
	)

	// Output:
	// Ace of Hearts
	// Two of Diamonds
	// King of Clubs
	// Queen of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards")
	}
}

func TestNewDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if len(cards) != 52 {
		t.Error("Wrong number of cards")
	}
}
func TestNewInverseDefaultSort(t *testing.T) {
	cards := New(InverseDefaultSort)
	if len(cards) != 52 {
		t.Error("Wrong number of cards")
	}
}
func TestCustomSort(t *testing.T) {
	cards := New(Sort(Higher))
	if len(cards) != 52 {
		t.Error("Wrong number of cards")
	}
}

func TestJoker(t *testing.T) {
	jokers := 5
	cards := New(Jokers(jokers))
	if len(cards) != 52+jokers {
		t.Error("Wrong number of jokers")
	}
}

func TestFilter(t *testing.T) {

	cards := New(
		Filter(
			func(card Card) bool {
				return card.Rank == Two
			},
		),
	)
	if len(cards) != 52-4 {
		t.Error("Wrong number of cards. ", len(cards))
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(4))
	if len(cards) != 52*4 {
		t.Error("Wrong number of cards")

	}
}
