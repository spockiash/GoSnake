package main

import (
	"strings"
)

// function that adds symbol to the grid
func AddSymbol(definition map[Coordinates]rune,
	position Coordinates,
	symbol rune) map[Coordinates]rune {
	definition[position] = symbol
	return definition
}

// simple function renders the arena defitinition into a string
func RenderDefinition(definition map[Coordinates]rune) string {
	height := arenaHeight()
	width := arenaWidth()

	lines := make([]string, height)

	// iterate over height (rows of lines)
	for h := 0; h < height; h++ {
		line := make([]rune, width)
		// iterate over width (line)
		for w := 0; w < width; w++ {
			position := Coordinates{
				h,
				w,
			}
			// use the coma ok idiom
			if r, ok := definition[position]; ok {
				line[w] = r
			} else {
				line[w] = ' ' // default to space
			}
		}
		lines[h] = string(line)
	}

	// construct output
	var sb strings.Builder
	for _, line := range lines {
		sb.WriteString(line)
		sb.WriteRune('\n')
	}
	return sb.String()
}
