package main

import (
	"github.com/rivo/tview"
)

type Arena struct {
	definition    map[Coordinates]rune
	renderedArena string
	arenaHeight   int
	arenaWidth    int
	arenaElement  *tview.TextView
	foodSource    []Food
}

type Coordinates struct {
	row int // height
	col int // width
}

// creates new arena, used when initalizing the game
func NewArena() Arena {
	// create map holding coordinates for optimal further write access
	definition := map[Coordinates]rune{}

	height := arenaHeight()
	width := arenaWidth()

	lines := make([]string, height)

	// iterate over height (rows of lines)
	for h := 0; h < height; h++ {
		line := make([]rune, width)
		// iterate over width (line)
		for w := 0; w < width; w++ {
			line[w] = ' '
			position := Coordinates{
				h,
				w,
			}
			definition[position] = line[w]
		}
		lines[h] = string(line)
	}

	renderedContent := AssembleString(lines)

	// hardcoded for now
	foodSource := []Food{
		{
			Coordinates{
				col: 5,
				row: 5,
			},
		},
		{
			Coordinates{
				col: 25,
				row: 30,
			},
		},
		{
			Coordinates{
				col: 15,
				row: 28,
			},
		},
	}

	return Arena{
		definition,
		renderedContent,
		height,
		width,
		createArenaElement(renderedContent),
		foodSource,
	}
}

// clears the arena definition (used before render cycles)
func ClearArena(arena *Arena) {
	arena.definition = map[Coordinates]rune{}
}

// creates the arena tview element
func createArenaElement(content string) *tview.TextView {
	arena := tview.NewTextView().
		SetText(content).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
	arena.SetBorder(true)
	return arena
}

func arenaWidth() int {
	return Width - 2
}

func arenaHeight() int {
	return Height - 2
}
