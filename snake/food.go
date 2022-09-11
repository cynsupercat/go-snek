package snake

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

const foodChar = '●'

var (
	foodStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorReset)
)

type Food struct {
	posX, posY int
	Eaten      bool
}

func NewFood() *Food {
	return &Food{
		posX: 0,
		posY: 0,
	}
}

func (food *Food) Draw(screen tcell.Screen) {
	screen.SetContent(food.posX, food.posY, foodChar, nil, foodStyle)
}

func (food *Food) Spawn(boundaryX, boundaryY int) {
	rand.Seed(time.Now().UnixNano())

	randPos := func(min, max int) int {
		return rand.Intn(max-min) + min
	}

	food.posY = randPos(0, boundaryY)
	food.posX = randPos(0, boundaryX)
}
