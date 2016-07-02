package blackjack

type Player struct {
	Hand []Card
	deck *Deck
}

func (p *Player) DrawCard() Card {
	c := p.deck.DrawCard()
	p.Hand = append(p.Hand, c)
	return c
}

func (p *Player) Total() int {
	t := 0
	for _, card := range p.Hand {
		t += card.Value
	}
	return t
}

func (p *Player) Bust() bool {
	return p.Total() > 21
}

func (p *Player) CheckBlackjack() bool {
	return ((p.Hand[0].Value == 11 && p.Hand[1].Value == 10) || (p.Hand[1].Value == 11 && p.Hand[0].Value == 10))
}

func NewPlayer(d *Deck) Player {
	return Player{deck: d}
}
