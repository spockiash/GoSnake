package main

type Snake struct {
	Body          []Coordinates
	Direction     Coordinates // use coordinates as direction, {0,1} = right
	NextDirection Coordinates
	Healthy       bool
}

func NewSnake(row int, col int) Snake {
	return Snake{
		Body: []Coordinates{
			{row: row, col: col},
		},
		Direction: Coordinates{row: 0, col: 0},
		Healthy:   true,
	}
}

func MoveSnake(snake *Snake, arena *Arena, grow bool) {
	snake.Direction = snake.NextDirection

	head := snake.Body[0] //start of the body slice is head

	// extract directions
	v := clampDirection(snake.Direction.row)
	h := clampDirection(snake.Direction.col)

	// get new head
	newHead := Coordinates{
		row: head.row + v,
		col: head.col + h,
	}

	// when collision happens mark snake as unhealthy
	snake.Healthy = healthCheck(newHead, arena, snake)

	// prepend the head to the body slince and store as new body
	newBody := append([]Coordinates{newHead}, snake.Body...)

	// when no grow signal delete last new body member
	if !grow {
		// select slice without last item
		newBody = newBody[:len(newBody)-1]
	}

	// asign the body
	snake.Body = newBody

}

// draws the snake onto the arena
func DrawSnake(snake *Snake, arena *Arena) {
	for _, v := range snake.Body {
		arena.definition[v] = '*'
	}
}

func healthCheck(newHead Coordinates, arena *Arena, snake *Snake) bool {
	// check collisions with the arena
	upperBorder := 0
	lowerBorder := arena.arenaWidth
	leftBorder := 0
	rightBorder := arena.arenaWidth

	if newHead.row == upperBorder || newHead.row == lowerBorder {
		return false
	}

	if newHead.col == leftBorder || newHead.col == rightBorder {
		return false
	}

	// TODO check with itself
	return true
}

func clampDirection(val int) int {
	if val < -1 {
		return -1
	}
	if val > 1 {
		return 1
	}
	return val
}
