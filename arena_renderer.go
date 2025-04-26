package main

import (
	"fmt"
	"strings"
)

// function that adds symbol to the grid
func AddSymbol(definition map[Coordinates]rune,
	position Coordinates,
	symbol rune) map[Coordinates]rune {
	definition[position] = symbol
	return definition
}

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

			if r, ok := definition[position]; ok {
				line[w] = r
			} else {
				// Here we inject the row number on second and third column
				if w == 1 || w == 2 {
					// convert row number to digit runes
					digits := []rune(fmt.Sprintf("%02d", h))
					line[1] = digits[0] // put first digit at column 1
					line[2] = digits[1] // put second digit at column 2
				} else {
					line[w] = ' ' // default to space
				}
			}
		}
		lines[h] = string(line)
	}

	// construct output
	return AssembleString(lines)
}

// use stringbuilder to assemble rendered lines int ouptut
func AssembleString(lines []string) string {
	var sb strings.Builder
	for i, line := range lines {
		sb.WriteString(line)
		if i != len(lines)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}
