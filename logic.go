package main

import (
	"errors"
)

func (h *Hand) count() int {
	counts := [7]int{}
	for _, c := range h.Cards {
		counts[c.SideA]++
		counts[c.SideB]++
	}

	maxIndex := 0
	maxVal := 0
	for i, v := range counts {
		if v >= maxVal {
			maxIndex = i
			maxVal = v
		}
	}
	return maxIndex
}

func (h *Hand) playDoublesGreaterThenOrEqualTo(max int) (int, bool) {
	for i, c := range h.Cards {
		if !c.CanPlay() {
			continue
		}
		for d := 6; d >= max; d-- {
			if c.SideA == d && c.SideB == d {
				println("Playing double greater than", max)
				return i, true
			}
		}
	}
	return 0, false
}

func (h *Hand) squareItOffWithWhatIHaveMostOf(mostOf int) (int, Side, bool) {

	if board.Left() != mostOf && board.Right() != mostOf {
		return 0, LeftSide, false
	}

	for i, c := range h.Cards {

		if !c.CanPlay() {
			continue
		}

		if c.SideA != mostOf && c.SideB != mostOf {
			continue
		}

		if board.Left() == mostOf {
			if board.Right() == c.SideA {
				println("Squaring it off with what I have the most of on RIGHT", i)
				return i, RightSide, true
			}

			if board.Right() == c.SideB {
				//c.Flip()
				println("Squaring it off with what I have the most of on RIGHT", i)
				return i, RightSide, true
			}
		}

		if board.Right() == mostOf {
			if board.Left() == c.SideA {
				//c.Flip()
				println("Squaring it off with what I have the most of on LEFT", i)
				return i, RightSide, true
			}
			if board.Left() == c.SideB {
				println("Squaring it off with what I have the most of on LEFT", i)
				return i, RightSide, true
			}
		}

	}
	return 0, LeftSide, false
}

func (h *Hand) squareItOffIfPossible() (int, bool) {
	for i, c := range h.Cards {
		if !c.CanPlay() {
			continue
		}

		// If we can play a single card on either side, do it...
		if (board.Left() == c.SideA && board.Right() == c.SideB) ||
			(board.Left() == c.SideB && board.Right() == c.SideA) {
			println("Squaring it off")
			return i, true
		}

	}
	return 0, false
}

func (h *Hand) playWhatWeHaveMostOf(mostOf int) (int, bool) {
	for i, c := range h.Cards {
		if !c.CanPlay() {
			continue
		}

		if c.SideA == mostOf {
			if board.Left() == c.SideB {
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
			if board.Right() == c.SideB {
				//flip()
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
		}

		if c.SideB == mostOf {
			if board.Left() == c.SideA {
				//flip()
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
			if board.Right() == c.SideA {
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
		}

	}
	return 0, false
}

func (h *Hand) getFirstPlayableCard() (int, bool) {
	for i, c := range h.Cards {
		if c.CanPlay() {
			println("Playing first playable card")
			return i, true
		}
	}
	return 0, false
}

func (h *Hand) determineBestPlay() (int, Side, error) {

	println("Determining best play")
	mostOf := h.count()
	println("You have the most of", mostOf)

	if index, ok := h.playDoublesGreaterThenOrEqualTo(4); ok {
		return index, LeftSide, nil
	}

	if index, side, ok := h.squareItOffWithWhatIHaveMostOf(mostOf); ok {
		return index, side, nil
	}

	if index, ok := h.squareItOffIfPossible(); ok {
		return index, LeftSide, nil
	}

	if index, ok := h.playWhatWeHaveMostOf(mostOf); ok {
		return index, LeftSide, nil
	}

	if index, ok := h.playDoublesGreaterThenOrEqualTo(0); ok {
		return index, LeftSide, nil
	}

	if index, ok := h.getFirstPlayableCard(); ok {
		return index, LeftSide, nil
	}

	return 0, AskMe, errors.New("No playable cards")
}
