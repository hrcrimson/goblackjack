package blackjack

var Faces = []Face{
	{"2", 2},
	{"3", 3},
	{"4", 4},
	{"5", 5},
	{"6", 6},
	{"7", 7},
	{"8", 8},
	{"9", 9},
	{"10", 10},
	{"J", 10},
	{"Q", 10},
	{"K", 10},
	{"A", 11},
}

var Suits = []string{
	"♠",
	"♥",
	"♦",
	"♣",
}

type Face struct {
	Face  string
	Value int
}

type Card struct {
	Face
	Suit string
}

func (c Card) String() string {
	return c.Face.Face + c.Suit
}

func NewCard(f Face, s string) Card {
	return Card{Face: f, Suit: s}
}
