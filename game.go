package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
)

const MAX_PLAYERS = 2

type Game struct {
	id      string
	Players [MAX_PLAYERS]*Hand
	Draw    *DrawPile
	Board   *Board
	Turn    int
}

type PlayResult int

const (
	CardPlayed PlayResult = iota
	Passed
	WonGame
)

// NewGame creates a new game including a new hydrated draw pile.
func NewGame() *Game {
	return &Game{
		id:    guid(),
		Draw:  NewDrawPile(),
		Board: &Board{},
		Turn:  0,
	}
}

// AddPlayer add a player to the game.  If there are already two players, it returns an error.
func (g *Game) AddPlayer(h *Hand) error {
	for i := 0; i < MAX_PLAYERS; i++ {
		if g.Players[i] == nil {
			g.Players[i] = h
			return nil
		}
	}
	return errors.New(fmt.Sprintf("This game has %d players already", MAX_PLAYERS))
}

// Play is the driver of the game experience.
func (g *Game) Play() {
	g.drawCards()
	for true {
		// @todo: Check for a locked board...
		g.display()

		if play := g.play(); play == WonGame {
			g.display()
			return
		}
	}
}

func guid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (g *Game) drawCards() {
	for i := 1; i < 8; i++ {
		for j := 0; j < MAX_PLAYERS; j++ {
			g.Players[j].Cards = append(g.Players[j].Cards, g.Draw.Pick())
		}
	}
}

func (g *Game) display() {
	println(g.id)
	g.Board.Print()
	for j := 0; j < MAX_PLAYERS; j++ {
		g.Players[j].Print()
	}
	fmt.Printf("Player [%s]'s Turn\n", g.Players[g.Turn].Name)
}

func (g *Game) play() PlayResult {
	defer func() {
		g.Turn++
		if g.Turn >= MAX_PLAYERS {
			g.Turn = 0
		}
	}()

	p := g.Players[g.Turn]

	_, c := p.Play(g.Board, g.Draw)
	if len(p.Cards) == 0 {
		fmt.Printf("Player [%s] WINS!!!\n", p.Name)
		return WonGame
	}

	if c == nil {
		fmt.Printf("Player [%s] could NOT play\n", p.Name)
		return Passed
	}

	fmt.Printf("Player [%s] just played: [%d|%d]\n", p.Name, c.SideA, c.SideB)
	return CardPlayed
}
