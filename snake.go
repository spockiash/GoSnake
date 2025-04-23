package main

type Snake struct {
	Body      []Coordinates
	Direction Coordinates // use coordinates as direction, {0,1} = right
}

func MoveSnake(snake *Snake, grow bool) {
	head := snake.Body[0] //start of the body slice is head
	// extract directions
	v := ClampDirection(snake.Direction.row)
	h := ClampDirection(snake.Direction.col)

	// get new head
	newHead := Coordinates{
		row: head.row + v,
		col: head.col + h,
	}

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

func ClampDirection(val int) int {
	if val < -1 {
		return -1
	}
	if val > 1 {
		return 1
	}
	return val
}
