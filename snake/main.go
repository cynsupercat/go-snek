package snake

func Collided(snake *Snake, food *Food) bool {
	return snake.head.x == food.posX &&
		snake.head.y == food.posY
}
