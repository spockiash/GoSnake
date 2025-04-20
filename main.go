package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	statusBar := createStatusBar()
	arena := createArena()

	// Use fixed size layout instead of Flex
	grid := tview.NewGrid().
		SetRows(1, 40). // top padding, arena height, bottom padding
		SetColumns(80). // left padding, arena width, right padding
		AddItem(statusBar, 0, 0, 1, 1, 1, 0, false).
		AddItem(arena, 1, 0, 1, 1, 0, 0, true)

	arena.SetRect(0, 0, 80, 40)
	// Display initial size information
	//x, y, width, height := arena.GetRect()

	// testing render function
	test := createTestRender(2, 2, 'W')

	arena.SetText(test)
	//arena.SetText(fmt.Sprintf("Arena size: %dx%d\nPosition: %d,%d", width, height, x, y))

	// Capture key events
	arena.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			x, y, width, height := arena.GetRect()
			arena.SetText(fmt.Sprintf("[green]Up arrow pressed\nArena size: %dx%d\nPosition: %d,%d", width, height, x, y))
		case tcell.KeyDown:
			x, y, width, height := arena.GetRect()
			arena.SetText(fmt.Sprintf("[red]Down arrow pressed\nArena size: %dx%d\nPosition: %d,%d", width, height, x, y))
		case tcell.KeyEscape, tcell.KeyCtrlC:
			app.Stop() // Exit on ESC or Ctrl+C
		}
		return event
	})

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func createArena() *tview.TextView {
	arena := tview.NewTextView().
		SetText("[green]Press ↑ or ↓ to see input").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	arena.SetBorder(true)
	return arena
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

// testing rendering of single symbol to arbitrary position
func createTestRender(row int, col int, symbol rune) string {
	height := ArenaHeight()
	width := ArenaWidth()

	lines := make([]string, height)

	// iterate over height (rows of lines)
	for h := 0; h < height; h++ {
		line := make([]rune, ArenaWidth())
		// iterate over width (line)
		for w := 0; w < width; w++ {
			// match to position
			if h == row && w == col {
				line[w] = symbol
			}
		}
		lines[h] = string(line)
	}

	// construct output
	var sb strings.Builder

	for i := range lines {
		sb.WriteString(lines[i])
		sb.WriteRune('\n')
	}

	return sb.String()
}

func createTestString() string {
	runes := make([]rune, ArenaWidth())

	var counter int
	for i := range runes {
		if counter > 9 {
			counter = 0
		}
		runes[i] = rune('A' + counter) // adds A, B, C, D

		counter++
	}

	return string(runes)
}
