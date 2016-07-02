package blackjack

import (
	"fmt"
	"github.com/hrcrimson/goinput"
)

type Round struct {
	d      Deck
	User   Player
	Dealer Player
}

func NewRound() Round {
	deck := NewDeck()
	return Round{d: deck, User: NewPlayer(&deck), Dealer: NewPlayer(&deck)}
}

func (r *Round) Play() bool {
	fmt.Println("Welcome to blackjack!")
	fmt.Println("You got a", r.User.DrawCard(), "and a", r.User.DrawCard())
	fmt.Println("Your total is", r.User.Total())

	r.Dealer.DrawCard()
	r.Dealer.DrawCard()
	fmt.Println("The dealer has a", r.Dealer.Hand[0], "and a hidden card")
	if !r.User.CheckBlackjack() {
		stay := false
		for !stay && !r.User.Bust() {
			if goinput.ReadBoolOptions("Would you like to \"hit\" or \"stay\"?", "hit", "stay") {
				fmt.Println("You drew a", r.User.DrawCard())
				fmt.Println("Your total is", r.User.Total())
			} else {
				stay = true
			}
		}
		if !r.User.Bust() {
			fmt.Println("Dealer's turn")
			fmt.Println("His hidden card was a", r.Dealer.Hand[1])
			fmt.Println("His total is", r.Dealer.Total())
			for r.Dealer.Total() < 16 {
				fmt.Println("Dealer hits")
				fmt.Println("He drew a", r.Dealer.DrawCard())
				fmt.Println("His total is", r.Dealer.Total())
			}
			if !r.Dealer.Bust() {
				fmt.Println("Dealer stays")
				fmt.Println("Dealer total:", r.Dealer.Total())
				fmt.Println("Your total", r.User.Total())
				if r.Dealer.Total() >= r.User.Total() {
					fmt.Println("Dealer wins")
					return false
				}
				fmt.Println("You win")
				return true
			}
			fmt.Println("Dealer busts! You win")
			return true
		}
		fmt.Println("Bust! You lose")
		return false

	}
	fmt.Println("Blackjack! You win!")
	return true

}
