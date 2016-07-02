package blackjack

import (
	"time"
	"math/rand"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) DrawCard() Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(d.Cards))
	c := d.Cards[i]
	d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
	return c
}

func NewDeck() Deck {
	d := Deck{}
	for _, suite := range Suits {
		for _, face := range Faces {
			d.Cards = append(d.Cards, NewCard(face, suite))
		}
	}
	return d
}