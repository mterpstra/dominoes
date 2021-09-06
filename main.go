package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Side int

const (
	LeftSide Side = iota
	RightSide
	AskMe
)

type DrawPile struct {
	Cards []*Card
}

var player1 Hand
var player2 Hand
var draw DrawPile
var board Board

var source rand.Source
var random *rand.Rand

func init() {
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			draw.Cards = append(draw.Cards, &Card{SideA: i, SideB: j})
		}
	}
	source = rand.NewSource(time.Now().UnixNano())
	random = rand.New(source)
}

func ChooseLeftOrRight() string {
	side := ""
	for side != "L" && side != "R" {
		fmt.Printf("Enter Side of Board (L,R):")
		fmt.Scanf("%s", &side)
		if side == "l" {
			side = "L"
		}
		if side == "r" {
			side = "R"
		}
	}
	return side
}

func ChooseNumberInRange(min, max int) int {
	value := 0
	valid := false
	for !valid {
		fmt.Printf("Enter Card Index (%d,%d):", min, max)
		fmt.Scanf("%d", &value)
		if value >= min && value <= max {
			valid = true
		}
	}
	return value
}

func (d *DrawPile) Pick() *Card {
	if len(d.Cards) == 0 {
		println("Draw pile is empty")
		return nil
	}

	i := random.Intn(len(d.Cards))
	card := d.Cards[i]
	d.Cards[i] = d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}

func (b *Board) Play(h *Hand, index int, side Side) (error, *Card) {

	if len(board.Cards) == 0 {
		card := h.Cards[index]
		h.Cards[index] = h.Cards[len(h.Cards)-1]
		h.Cards = h.Cards[:len(h.Cards)-1]
		board.Cards = append([]*Card{card}, board.Cards...)
		return nil, card
	}

	playOnLeft := h.Cards[index].SideA == board.Left() || h.Cards[index].SideB == board.Left()
	playOnRight := h.Cards[index].SideA == board.Right() || h.Cards[index].SideB == board.Right()

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

	if playOnLeft && h.Cards[index].SideB != board.Left() {
		h.Cards[index].Flip()
	}

	if playOnRight && h.Cards[index].SideA != board.Right() {
		h.Cards[index].Flip()
	}

	card := h.Cards[index]
	h.Cards[index] = h.Cards[len(h.Cards)-1]
	h.Cards = h.Cards[:len(h.Cards)-1]
	if playOnLeft {
		board.Cards = append([]*Card{card}, board.Cards...)
	} else {
		board.Cards = append(board.Cards, card)
	}

	return nil, card
}

func (h *Hand) CanPlay() bool {
	for _, c := range h.Cards {
		if c.CanPlay() {
			return true
		}
	}
	return false
}

func (h *Hand) Play() (error, *Card) {
	if h.IsComputer {
		return h.PlayAuto()
	} else {
		return h.PlayManual()
	}
}

func (h *Hand) PlayAuto() (error, *Card) {

	for !h.CanPlay() {
		println("can't play, drawing...")
		if c := draw.Pick(); c != nil {
			h.Cards = append(h.Cards, c)
			h.Print()
		} else {
			// @todo: Should can't play be an error?
			return nil, nil
		}
	}

	index, side, _ := h.determineBestPlay()
	return board.Play(h, index, side)
}

func (h *Hand) PlayManual() (error, *Card) {
	for !h.CanPlay() {
		var u string
		println("You can't play, draw from pile (d)")
		fmt.Scanf("%s", &u)

		if c := draw.Pick(); c != nil {
			h.Cards = append(h.Cards, c)
			h.Print()
		} else {
			return nil, nil
		}
	}

	index := 0
	for true {
		index = ChooseNumberInRange(0, len(h.Cards)-1)
		if h.Cards[index].CanPlay() {
			break
		}
		println("That card cannot be played")
	}

	return board.Play(h, index, AskMe)
}

func printCards(desc string, c []*Card) {
	fmt.Printf("%s", desc)
	for _, v := range c {
		fmt.Printf("[%d|%d] ", v.SideA, v.SideB)
	}
	fmt.Printf("  (len=%d)\n", len(c))
}

func printDrawPile() {
	printCards("\033[0;31mdraw    ", draw.Cards)
	print("\033[0m") // Reset
}

func main() {

	player1 = Hand{
		Name:       "Computer",
		IsComputer: true,
		Color:      "\033[0;32m",
	}

	player2 = Hand{
		Name:       "Mark",
		IsComputer: true,
		Color:      "\033[0;33m",
	}

	for i := 1; i < 8; i++ {
		player1.Cards = append(player1.Cards, draw.Pick())
		player2.Cards = append(player2.Cards, draw.Pick())
	}

	turn := true
	msg := "Starting Game"
	for true {
		// @todo: Check for a locked board...

		// Print the board
		board.Print()
		player1.Print()
		player2.Print()
		fmt.Printf("Last Play: %s\n", msg)

		if turn {
			fmt.Printf("Player 1's Turn\n")
		} else {
			fmt.Printf("Player 2's Turn\n")
		}

		if turn {
			_, card := player1.Play()
			if len(player1.Cards) == 0 {
				println("player1 Wins!")
				break
			}

			if card == nil {
				msg = fmt.Sprintf("Player 1 could NOT play")
			} else {
				msg = fmt.Sprintf("Player 1 just played: [%d|%d]\n", card.SideA, card.SideB)
			}

		}

		if !turn {
			_, card := player2.Play()
			if len(player2.Cards) == 0 {
				println("player2 Wins!")
				break
			}
			if card == nil {
				msg = fmt.Sprintf("Player 2 could NOT play")
			} else {
				msg = fmt.Sprintf("Player 2 just played: [%d|%d]\n", card.SideA, card.SideB)
			}
		}

		turn = !turn
	}
	println("End of Game")
}
