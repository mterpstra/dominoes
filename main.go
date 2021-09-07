package main

func main() {

	game := NewGame()
	game.AddPlayer(&Hand{
		Name:       "Computer",
		IsComputer: true,
		Color:      "\033[0;32m",
	})

	game.AddPlayer(&Hand{
		Name:       "Mark",
		IsComputer: false,
		Color:      "\033[0;33m",
	})

	game.DrawCards()
	game.Play()
}
