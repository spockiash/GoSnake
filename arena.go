package main

import (
	"math/rand"

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

	foodSource := generateFood()

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

// draws the food onto the arena
func DrawFood(arena *Arena) {
	for _, food := range arena.foodSource {
		arena.definition[food.Position] = '@'
	}
}

func generateFood() []Food {
	amount := (Height * Height) / 12
	food := []Food{}
	for i := 0; i < amount; i++ {
		col := rand.Intn(arenaWidth()) + 1
		row := rand.Intn(arenaHeight()) + 1
		food = append(food, Food{
			Position: Coordinates{
				row: row,
				col: col,
			},
		})
	}

	return food
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
