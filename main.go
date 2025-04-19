package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create the arena with fixed dimensions
	arena := tview.NewTextView().
		SetText("[green]Press ↑ or ↓ to see input").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	arena.SetBorder(true)

	// Use fixed size layout instead of Flex
	grid := tview.NewGrid().
		SetRows(0, 40, 0).    // top padding, arena height, bottom padding
		SetColumns(0, 80, 0). // left padding, arena width, right padding
		AddItem(arena, 1, 1, 1, 1, 0, 0, true)

	// Display initial size information
	x, y, width, height := arena.GetRect()
	arena.SetText(fmt.Sprintf("Arena size: %dx%d\nPosition: %d,%d", width, height, x, y))

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
