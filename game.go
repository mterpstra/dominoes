package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os"
)

// @todo: Support more than 2 players
type Game struct {
	id            string
	Player1       *Hand
	Player2       *Hand
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
	if g.Player1 == nil {
		g.Player1 = h
	} else if g.Player2 == nil {
		g.Player2 = h
	} else {
		return errors.New("This game has 2 players already")
	}
	return nil
}

// DrawCards is called when the game begins and each player
// draws 7 cards at random.
func (g *Game) DrawCards() {
	for i := 1; i < 8; i++ {
		g.Player1.Cards = append(g.Player1.Cards, g.Draw.Pick())
		g.Player2.Cards = append(g.Player2.Cards, g.Draw.Pick())
	}
}

// Print is used to print the game.
func (g *Game) display() {
	println(g.id)
	g.Board.Print()
	g.Player1.Print()
	g.Player2.Print()

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
		_, card := g.Player1.Play(g.Board, g.Draw)
		if len(g.Player1.Cards) == 0 {
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
		_, card := g.Player2.Play(g.Board, g.Draw)
		if len(g.Player2.Cards) == 0 {
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
