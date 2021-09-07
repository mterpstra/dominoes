package main

type Side int

const (
	LeftSide Side = iota
	RightSide
	AskMe
)

type Card struct {
	SideA    int
	SideB    int
	Playable bool
}

func (c *Card) Flip() {
	tmp := c.SideA
	c.SideA = c.SideB
	c.SideB = tmp
}

func (c *Card) CanPlay(b *Board) bool {
	// The board is empty, so yes this card can be played
	if len(b.Cards) == 0 {
		return true
	}

	// Check the card on both ends
	return c.SideA == b.Left() ||
		c.SideB == b.Left() ||
		c.SideA == b.Right() ||
		c.SideB == b.Right()
}
