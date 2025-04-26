package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	statusBar := createStatusBar()
	arena := NewArena()
	snake := NewSnake((Height/2)-2, Width/2) // set initial position to middle of the arena, -2 is border offset

	// initial snake draw
	DrawSnake(&snake, &arena)
	arena.renderedArena = RenderDefinition(arena.definition)
	arena.arenaElement.SetText(arena.renderedArena)

	// fixed size layout
	grid := tview.NewGrid().
		SetRows(1, Height). // status bar height and arena height
		SetColumns(Width).
		AddItem(statusBar, 0, 0, 1, 1, 0, 0, false).
		AddItem(arena.arenaElement, 1, 0, 1, 1, 0, 0, true)

	// channel to signal shutdown
	done := make(chan struct{})

	// input handler
	arena.arenaElement.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if snake.Direction.row != 1 {
				snake.NextDirection = Coordinates{row: -1, col: 0}
			}
		case tcell.KeyDown:
			if snake.Direction.row != -1 {
				snake.NextDirection = Coordinates{row: 1, col: 0}
			}
		case tcell.KeyLeft:
			if snake.Direction.col != 1 {
				snake.NextDirection = Coordinates{row: 0, col: -1}
			}
		case tcell.KeyRight:
			if snake.Direction.col != -1 {
				snake.NextDirection = Coordinates{row: 0, col: 1}
			}
		case tcell.KeyEscape, tcell.KeyCtrlC:
			close(done)
			app.Stop()
		}
		return event
	})

	// proper game loop using ticker
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		// helper variable to end the game loop
		stopGameLoop := false

		for {
			select {
			case <-ticker.C:
				app.QueueUpdateDraw(func() {
					// check for health
					if snake.Healthy != true {
						arena.arenaElement.SetText("[red]Game Over")
						stopGameLoop = true
						return
					}
					// perform movement logic, this updates snake body positions
					MoveSnake(&snake, &arena, false)

					// clear game arena before render pass
					ClearArena(&arena)

					// adds updated snake to definition
					DrawSnake(&snake, &arena)

					// perform the render pass
					arena.renderedArena = RenderDefinition(arena.definition)
					arena.arenaElement.SetText(arena.renderedArena)
				})
			case <-done:
				return
			}
			// break out of the loop
			if stopGameLoop {
				break
			}
		}
	}()

	// run application
	if err := app.SetRoot(grid, true).EnableMouse(false).Run(); err != nil {
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
