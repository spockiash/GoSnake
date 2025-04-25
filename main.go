package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	statusBar := createStatusBar()
	arena := NewArena()
	snake := NewSnake((Height/2)-2, Width/2) // set initial position to middle of the arena, -2 is border offset

	// draw snake onto the arena definition
	DrawSnake(&snake, &arena)

	// render the arena based on definition
	arena.renderedArena = RenderDefinition(arena.definition)
	arena.arenaElement.SetText(arena.renderedArena)

	// Use fixed size layout instead of Flex
	grid := tview.NewGrid().
		SetRows(1, 40). // top padding, arena height, bottom padding
		SetColumns(80). // left padding, arena width, right padding
		AddItem(statusBar, 0, 0, 1, 1, 1, 0, false).
		AddItem(arena.arenaElement, 1, 0, 1, 1, 0, 0, true)

	arena.arenaElement.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:

		case tcell.KeyDown:

		case tcell.KeyEscape, tcell.KeyCtrlC:
			app.Stop() // Exit on ESC or Ctrl+C
		}
		return event
	})

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func createStatusBar() *tview.Flex {
	score := tview.NewTextView().
		SetText("Score: 420").
		SetTextAlign(tview.AlignLeft)

	time := tview.NewTextView().
		SetText("4:20").
		SetTextAlign(tview.AlignRight)

	flex := tview.NewFlex().
		AddItem(score, 0, 1, false).
		AddItem(time, 0, 1, false)

	return flex
}

// // testing rendering of single symbol to arbitrary position
// func createTestRender(row int, col int, symbol rune) (string, map[Coordinates]rune) {
// 	// TODO: boundary check (right now if out of bounds, the condition will simply
// 	// not be satisfied and empty render will be created)

// 	// create map holding coordinates for optimal further write access
// 	lookup := map[Coordinates]rune{}

// 	height := ArenaHeight()
// 	width := ArenaWidth()

// 	lines := make([]string, height)

// 	// iterate over height (rows of lines)
// 	for h := 0; h < height; h++ {
// 		line := make([]rune, ArenaWidth())
// 		// iterate over width (line)
// 		for w := 0; w < width; w++ {
// 			// match to position
// 			if h == row && w == col {
// 				line[w] = symbol
// 			} else {
// 				line[w] = ' '
// 				continue
// 			}
// 		}
// 		lines[h] = string(line)
// 	}

// 	// construct output
// 	var sb strings.Builder
// 	for _, line := range lines {
// 		sb.WriteString(line)
// 		sb.WriteRune('\n')
// 	}

// 	return sb.String(), lookup
// }

// func createTestString() string {
// 	runes := make([]rune, ArenaWidth())

// 	var counter int
// 	for i := range runes {
// 		if counter > 9 {
// 			counter = 0
// 		}
// 		runes[i] = rune('A' + counter) // adds A, B, C, D

// 		counter++
// 	}

// 	return string(runes)
// }
