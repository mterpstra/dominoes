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

func (h *Hand) playDoublesGreaterThenOrEqualTo(b *Board, max int) (int, bool) {
	for i, c := range h.Cards {
		if !c.CanPlay(b) {
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

func (h *Hand) squareItOffWithWhatIHaveMostOf(b *Board, mostOf int) (int, Side, bool) {

	if b.Left() != mostOf && b.Right() != mostOf {
		return 0, LeftSide, false
	}

	for i, c := range h.Cards {

		if !c.CanPlay(b) {
			continue
		}

		if c.SideA != mostOf && c.SideB != mostOf {
			continue
		}

		if b.Left() == mostOf {
			if b.Right() == c.SideA {
				println("Squaring it off with what I have the most of on RIGHT", i)
				return i, RightSide, true
			}

			if b.Right() == c.SideB {
				//c.Flip()
				println("Squaring it off with what I have the most of on RIGHT", i)
				return i, RightSide, true
			}
		}

		if b.Right() == mostOf {
			if b.Left() == c.SideA {
				//c.Flip()
				println("Squaring it off with what I have the most of on LEFT", i)
				return i, RightSide, true
			}
			if b.Left() == c.SideB {
				println("Squaring it off with what I have the most of on LEFT", i)
				return i, RightSide, true
			}
		}

	}
	return 0, LeftSide, false
}

func (h *Hand) squareItOffIfPossible(b *Board) (int, bool) {
	for i, c := range h.Cards {
		if !c.CanPlay(b) {
			continue
		}

		// If we can play a single card on either side, do it...
		if (b.Left() == c.SideA && b.Right() == c.SideB) ||
			(b.Left() == c.SideB && b.Right() == c.SideA) {
			println("Squaring it off")
			return i, true
		}

	}
	return 0, false
}

func (h *Hand) playWhatWeHaveMostOf(b *Board, mostOf int) (int, bool) {
	for i, c := range h.Cards {
		if !c.CanPlay(b) {
			continue
		}

		if c.SideA == mostOf {
			if b.Left() == c.SideB {
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
			if b.Right() == c.SideB {
				//flip()
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
		}

		if c.SideB == mostOf {
			if b.Left() == c.SideA {
				//flip()
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
			if b.Right() == c.SideA {
				println("Playing what I have most of, shooting those", mostOf)
				return i, true
			}
		}

	}
	return 0, false
}

func (h *Hand) getFirstPlayableCard(b *Board) (int, bool) {
	for i, c := range h.Cards {
		if c.CanPlay(b) {
			println("Playing first playable card")
			return i, true
		}
	}
	return 0, false
}

func (h *Hand) determineBestPlay(b *Board) (int, Side, error) {

	println("Determining best play")
	mostOf := h.count()
	println("You have the most of", mostOf)

	if index, ok := h.playDoublesGreaterThenOrEqualTo(b, 4); ok {
		return index, LeftSide, nil
	}

	if index, side, ok := h.squareItOffWithWhatIHaveMostOf(b, mostOf); ok {
		return index, side, nil
	}

	if index, ok := h.squareItOffIfPossible(b); ok {
		return index, LeftSide, nil
	}

	if index, ok := h.playWhatWeHaveMostOf(b, mostOf); ok {
		return index, LeftSide, nil
	}

	if index, ok := h.playDoublesGreaterThenOrEqualTo(b, 0); ok {
		return index, LeftSide, nil
	}

	if index, ok := h.getFirstPlayableCard(b); ok {
		return index, LeftSide, nil
	}

	return 0, AskMe, errors.New("No playable cards")
}
