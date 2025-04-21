package main

import (
	"strings"
)

type Arena struct {
	definition    map[Coordinates]rune
	renderedArena string
	arenaHeight   int
	arenaWidth    int
}

// testing rendering of single symbol to arbitrary position
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
		}
		lines[h] = string(line)
	}

	// construct output
	var sb strings.Builder
	for _, line := range lines {
		sb.WriteString(line)
		sb.WriteRune('\n')
	}

	return Arena{
		definition,
		sb.String(),
		height,
		width,
	}
}

func arenaWidth() int {
	return Width - 2
}

func arenaHeight() int {
	return Height - 2
}
