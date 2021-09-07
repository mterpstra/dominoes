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

// Check the card on both ends
func (c *Card) CanPlay(b *Board) bool {

	if len(b.Cards) == 0 {
		return true
	}

	return c.SideA == b.Left() ||
		c.SideB == b.Left() ||
		c.SideA == b.Right() ||
		c.SideB == b.Right()
}
