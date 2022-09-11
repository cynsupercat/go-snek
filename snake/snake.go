package snake

import "github.com/gdamore/tcell/v2"

type axis struct {
	x, y int
}

type Snake struct {
	len        int
	tail, head axis
}

const (
	snakeChar = '■'
)

var (
	snakeStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorReset)
	defStyle   = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
)

func NewSnake(screen tcell.Screen) *Snake {
	width, height := screen.Size()

	defAxis := axis{
		x: width / 2,
		y: height / 2,
	}

	return &Snake{
		len:  1,
		head: defAxis,
		tail: defAxis,
	}
}

func (snake *Snake) Grow() {
	snake.len++
}

func (snake *Snake) Draw(screen tcell.Screen) {
	// if snake.len == 1 {
	// 	screen.SetContent(snake.head.x, snake.head.y, snakeChar, nil, snakeStyle)
	// 	return
	// }
	screen.SetContent(snake.head.x, snake.head.y, snakeChar, nil, snakeStyle)
	return
	// TODO logic shit. For now pretend its always a straight line
	// Linear logic
	// for i := 1; i <= snake.len; i++ {
	// 	screen.SetContent(snake.head.x, snake.head.x, snakeChar, nil, snakeStyle)
	// }
}

// TODO for now clear everything. Optimise by only clearing needed cells later
func (snake *Snake) clear(screen tcell.Screen) {
	screen.SetContent(snake.head.x, snake.head.y, ' ', nil, defStyle)
}

func (snake *Snake) Right(screen tcell.Screen) {
	snake.clear(screen)
	snake.head.x += 1
	snake.tail.x += 1
	snake.Draw(screen)
}

func (snake *Snake) Left(screen tcell.Screen) {
	snake.clear(screen)
	snake.tail.x -= 1
	snake.head.x -= 1
	snake.Draw(screen)
}

func (snake *Snake) Up(screen tcell.Screen) {
	snake.clear(screen)
	snake.tail.y += 1
	snake.head.y += 1
	snake.Draw(screen)
}

func (snake *Snake) Down(screen tcell.Screen) {
	snake.clear(screen)
	snake.tail.y -= 1
	snake.head.y -= 1
	snake.Draw(screen)
}
