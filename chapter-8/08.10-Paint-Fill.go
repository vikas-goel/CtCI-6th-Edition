package main

import "fmt"

type Color int
type Point struct {
	row, col int
}

const Black = 1
const Green = 2
const Red = 3
const Yellow = 4
const White = 5

func PaintFill(screen [][]Color, p Point, c Color) {
	// Invalid coordinate.
	if len(screen) < p.row-1 || len(screen[0]) < p.col-1 {
		return
	}

	// No change in color.
	if screen[p.row][p.col] == c {
		return
	}

	original := screen[p.row][p.col]

	target := make([]*Point, 0, len(screen)*len(screen[0]))
	target = append(target, &p)
	screen[p.row][p.col] = c

	for i, targetLen := 0, len(target); i < targetLen; i++ {
		lt := &Point{target[i].row-1, target[i].col-1}
		rt := &Point{target[i].row+1, target[i].col-1}
		rb := &Point{target[i].row+1, target[i].col+1}
		lb := &Point{target[i].row-1, target[i].col+1}
		st := &Point{target[i].row, target[i].col-1}
		rs := &Point{target[i].row+1, target[i].col}
		sb := &Point{target[i].row, target[i].col+1}
		ls := &Point{target[i].row-1, target[i].col}

		// Scan all 8 neighbor cells. Add each valid cell, that has
		// the original color of the original point, to the target.
		for _, pt := range []*Point{lt, rt, rb, lb, st, rs, sb, ls} {
			if pt.row < 0 || pt.row >= len(screen) ||
			   pt.col < 0 || pt.col >= len(screen[pt.row]) ||
		           screen[pt.row][pt.col] != original {
				continue
			}

			// Change color of the current screen.
			screen[pt.row][pt.col] = c
			target = append(target, pt)
		}

		targetLen = len(target)
	}
}

func main() {
	screen := [][]Color{
		{Black, White, White, Red, Red},
		{Green, Yellow, Red, Black, Red},
		{White, Yellow, Red, Green, Black},
		{Red, Black, Green, Red, White},
		{Green, Black, Green, Yellow, White} }

	fmt.Println(screen)
	PaintFill(screen, Point{2, 2}, Black)
	fmt.Println(screen)
}
