package main

import (
	"errors"
	"fmt"
)

type Board struct {
	Cards []*Card
}

func (b *Board) Left() int {
	if len(b.Cards) == 0 {
		return -1
	}
	return b.Cards[0].SideA
}

func (b *Board) Right() int {
	if len(b.Cards) == 0 {
		return -1
	}
	return b.Cards[len(b.Cards)-1].SideB
}

func (b *Board) Print() {
	print("\033[0;34m")
	defer func() {
		fmt.Printf("\n")
		print("\033[0m") // Reset
	}()

	if len(b.Cards) == 0 {
		fmt.Printf("Empty Board\n")
		return
	}

	for i := 0; i < len(b.Cards); i++ {
		fmt.Printf("┌───┬───┐")
	}
	fmt.Printf("\n")

	for _, v := range b.Cards {
		fmt.Printf("│ %d │ %d │", v.SideA, v.SideB)
	}
	fmt.Printf("\n")

	for i := 0; i < len(b.Cards); i++ {
		fmt.Printf("└───┴───┘")
	}
}

func (b *Board) Play(h *Hand, index int, side Side) (error, *Card) {

	if len(b.Cards) == 0 {
		card := h.Cards[index]
		h.Cards[index] = h.Cards[len(h.Cards)-1]
		h.Cards = h.Cards[:len(h.Cards)-1]
		b.Cards = append([]*Card{card}, b.Cards...)
		return nil, card
	}

	playOnLeft := h.Cards[index].SideA == b.Left() || h.Cards[index].SideB == b.Left()
	playOnRight := h.Cards[index].SideA == b.Right() || h.Cards[index].SideB == b.Right()

	if !playOnLeft && !playOnRight {
		return errors.New("Card cannot be played"), nil
	}

	if playOnLeft && playOnRight {
		if h.IsComputer {
			if side == LeftSide {
				playOnRight = false
			} else {
				playOnLeft = false
			}
		} else if side := ChooseLeftOrRight(); side == "L" {
			playOnRight = false
		} else {
			playOnLeft = false
		}
	}

	if playOnLeft && h.Cards[index].SideB != b.Left() {
		h.Cards[index].Flip()
	}

	if playOnRight && h.Cards[index].SideA != b.Right() {
		h.Cards[index].Flip()
	}

	card := h.Cards[index]
	h.Cards[index] = h.Cards[len(h.Cards)-1]
	h.Cards = h.Cards[:len(h.Cards)-1]
	if playOnLeft {
		b.Cards = append([]*Card{card}, b.Cards...)
	} else {
		b.Cards = append(b.Cards, card)
	}

	return nil, card
}

func (b *Board) CanPlay(c *Card) bool {
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
