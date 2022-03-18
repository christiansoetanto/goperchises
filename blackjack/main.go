package main

import (
	"fmt"
	"gophercises/deck"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {

	return h[0].String() + " , **HIDDEN**"

}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}

	for _, card := range h {
		if card.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, card := range h {
		score += min(int(card.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card

	var dealer Hand
	const playerNumber = 3
	var players [playerNumber]Hand

	for i := 0; i < 2; i++ {

		card, cards = draw(cards)
		dealer = append(dealer, card)

		for i := 0; i < playerNumber; i++ {
			player := &players[i]
			card, cards = draw(cards)
			*player = append(*player, card)
		}
	}

	for i := 0; i < playerNumber; i++ {
		player := &players[i]
		var input string
		for input != "s" {
			fmt.Println("===============================")

			fmt.Println("Player ", i+1)
			fmt.Println("Player card: ", player)
			fmt.Println("Dealer card: ", dealer.DealerString())
			fmt.Println("What are you gonna do? (h)it or (s)tand?")
			_, err := fmt.Scanf("%s\n", &input)
			if err != nil {
				return
			}

			switch input {
			case "h":
				{
					card, cards = draw(cards)
					*player = append(*player, card)
				}

			}

		}
	}

	dealerScore, dealerMinScore := dealer.Score(), dealer.MinScore()

	for dealerScore <= 16 || dealerScore == 17 && dealerMinScore != 17 {
		card, cards = draw(cards)
		dealer = append(dealer, card)
		dealerScore, dealerMinScore = dealer.Score(), dealer.MinScore()
	}
	fmt.Println("===============================")

	fmt.Println("Game ended. Result: ")
	fmt.Println("===============================")
	fmt.Println("Dealer card: ", dealer)
	fmt.Println("Dealer score: ", dealerScore)
	fmt.Println("===============================")

	for i, player := range players {

		playerScore := player.Score()
		fmt.Println("Player ", i+1)
		fmt.Println("Card: ", player)
		fmt.Println("Score: ", playerScore)
		switch {

		case playerScore > 21:
			fmt.Println("You busted")
		case dealerScore > 21:
			fmt.Println("Dealer busted")
		case playerScore > dealerScore:
			fmt.Println("You win")
		case playerScore < dealerScore:
			fmt.Println("You lose")
		case playerScore == dealerScore:
			fmt.Println("Tie")

		}

		fmt.Println("===============================")
	}

}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
