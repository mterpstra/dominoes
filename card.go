package main

type Card struct {
	SideA    int
	SideB    int
	Playable bool
}

func (c *Card) CanPlay() bool {
	// The board is empty, so yes this card can be played
	if len(board.Cards) == 0 {
		return true
	}

	// Check the card on both ends
	return c.SideA == board.Left() ||
		c.SideB == board.Left() ||
		c.SideA == board.Right() ||
		c.SideB == board.Right()
}

func (c *Card) Flip() {
	tmp := c.SideA
	c.SideA = c.SideB
	c.SideB = tmp
}
