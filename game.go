package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os"
)

const MAX_PLAYERS = 2

// @todo: Support more than 2 players
type Game struct {
	id            string
	Players       [MAX_PLAYERS]*Hand
	Draw          *DrawPile
	Board         *Board
	PlayerOneTurn bool
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

func NewGame() *Game {
	return &Game{
		id:            guid(),
		Draw:          NewDrawPile(),
		Board:         &Board{},
		PlayerOneTurn: true,
	}
}

// AddPlayer add a player to the game.  If there
// are already two players, it returns an error.
func (g *Game) AddPlayer(h *Hand) error {

	for i := 0; i < MAX_PLAYERS; i++ {
		if g.Players[i] == nil {
			g.Players[i] = h
			return nil
		}
	}

	return errors.New(fmt.Sprintf("This game has %d players already", MAX_PLAYERS))
}

// DrawCards is called when the game begins and each player
// draws 7 cards at random.
func (g *Game) DrawCards() {
	for i := 1; i < 8; i++ {
		for j := 0; j < MAX_PLAYERS; j++ {
			g.Players[j].Cards = append(g.Players[j].Cards, g.Draw.Pick())
		}
	}
}

// Print is used to print the game.
func (g *Game) display() {
	println(g.id)
	g.Board.Print()
	for j := 0; j < MAX_PLAYERS; j++ {
		g.Players[j].Print()
	}

	if g.PlayerOneTurn {
		fmt.Printf("Player 1's Turn\n")
	} else {
		fmt.Printf("Player 2's Turn\n")
	}
}

func (g *Game) Play() {
	for true {
		// @todo: Check for a locked board...
		g.display()
		g.play()
	}
}

func (g *Game) play() {

	msg := ""
	defer func() {
		fmt.Printf("Last Play: %s\n", msg)
		g.PlayerOneTurn = !g.PlayerOneTurn
	}()

	if g.PlayerOneTurn {
		_, card := g.Players[0].Play(g.Board, g.Draw)
		if len(g.Players[0].Cards) == 0 {
			println("player1 Wins!")
			// @todo: Better return code
			os.Exit(0)
		}

		if card == nil {
			msg = fmt.Sprintf("Player 1 could NOT play")
		} else {
			msg = fmt.Sprintf("Player 1 just played: [%d|%d]\n", card.SideA, card.SideB)
		}
	}

	if !g.PlayerOneTurn {
		_, card := g.Players[1].Play(g.Board, g.Draw)
		if len(g.Players[1].Cards) == 0 {
			println("player2 Wins!")
			// @todo: Better return code
			os.Exit(0)
		}
		if card == nil {
			msg = fmt.Sprintf("Player 2 could NOT play")
		} else {
			msg = fmt.Sprintf("Player 2 just played: [%d|%d]\n", card.SideA, card.SideB)
		}
	}
}
