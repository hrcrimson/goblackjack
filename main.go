package main

import (
	"github.com/hrcrimson/goblackjack/blackjack"
	"fmt"
	"github.com/hrcrimson/goinput"
)

func main() {
	wins := 0
	losses := 0
	stop := false

	fmt.Println("Welcome to Blackjack!")
	for !stop {
		fmt.Println("Wins:",wins,"Losses:",losses)
		round := blackjack.NewRound()
		if round.Play() {
			wins++
		} else {
			losses++
		}
		stop = !goinput.ReadBoolOptions("Would you like to \"play\" or \"quit\"?", "play", "quit")

	}

}