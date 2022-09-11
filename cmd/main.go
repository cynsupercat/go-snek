package main

import (
	"log"
	"os"

	snakeGame "github.com/cynsupercat/go-snek/snake"
	"github.com/gdamore/tcell/v2"
)

func main() {
	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	screen.SetStyle(defStyle)

	// Clear screen
	screen.Clear()

	quit := func() {
		screen.Fini()
		os.Exit(0)
	}

	width, height := screen.Size()

	snake := snakeGame.NewSnake(screen)
	food := snakeGame.NewFood()

	food.Spawn(width, height)
	food.Draw(screen)
	snake.Draw(screen)

	for {
		screen.Show()

		event := screen.PollEvent()

		switch ev := event.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyRight {
				snake.Right(screen)
			} else if ev.Key() == tcell.KeyLeft {
				snake.Left(screen)
			} else if ev.Key() == tcell.KeyDown { // TODO these 2 are inverted somehow?
				snake.Up(screen)
			} else if ev.Key() == tcell.KeyUp {
				snake.Down(screen)
			}

			if snakeGame.Collided(snake, food) {
				snake.Grow()
				food.Spawn(width, height)
				food.Draw(screen)
			}
		}
	}
}
