package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Initialize screen
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v\n", err)
		os.Exit(1)
	}
	defer screen.Fini()

	// Clear screen
	screen.Clear()

	// Display instructions
	drawText(screen, 1, 1, "Press arrow keys (up/down) to see output. Press ESC to quit.", tcell.StyleDefault)

	// Current line for output
	currentLine := 3

	// Display screen
	screen.Show()

	// Main event loop
	for {
		// Poll for events
		event := screen.PollEvent()

		// Process key events
		switch ev := event.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyUp:
				// Draw the message on screen
				drawText(screen, 1, currentLine, "Key up pressed", tcell.StyleDefault)
				currentLine++
			case tcell.KeyDown:
				// Draw the message on screen
				drawText(screen, 1, currentLine, "Key down pressed", tcell.StyleDefault)
				currentLine++
			case tcell.KeyEscape:
				return // Exit the program
			}
			// Update the screen after drawing new text
			screen.Show()
		}
	}
}

// Helper function to draw text on the screen
func drawText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}
