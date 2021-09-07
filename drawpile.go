package main

import (
	"math/rand"
	"time"
)

type DrawPile struct {
	Cards []*Card
}

func NewDrawPile() *DrawPile {
	d := DrawPile{}
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d.Cards = append(d.Cards, &Card{SideA: i, SideB: j})
		}
	}
	return &d
}

func (d *DrawPile) Pick() *Card {
	if len(d.Cards) == 0 {
		println("Draw pile is empty")
		return nil
	}

	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Intn(len(d.Cards))
	card := d.Cards[i]
	d.Cards[i] = d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}
