package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	f := Parse(input)

	log.Println("Trees visible:", f.visibleTrees())
	log.Println("Highest score:", f.highestScore())
}

func Parse(input string) Forest {
	var forest [][]int

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		var f []int
		for _, c := range row {
			n, _ := strconv.Atoi(string(c))
			f = append(f, n)
		}
		forest = append(forest, f)
	}

	return forest
}

type Forest [][]int

func (f Forest) visibleTrees() int {
	count := 0
	for y, r := range f {
		for x, h := range r {
			if f.isVisible(x, y, h) {
				count++
			}

		}
	}
	return count
}

func (f Forest) highestScore() int {
	highest := 0
	for y, r := range f {
		for x, h := range r {
			score := f.score(x, y, h)
			if score > highest {
				highest = score
			}
		}
	}
	return highest
}

func (f Forest) score(x, y, h int) int {
	score := f.visibilityHorizontally(x, y, 1, h) *
		f.visibilityHorizontally(x, y, -1, h) *
		f.visibilityVertically(x, y, 1, h) *
		f.visibilityVertically(x, y, -1, h)
	return score
}

func (f Forest) isVisible(x, y, h int) bool {

	return f.isVisibleVertically(x, y, 1, h) || f.isVisibleVertically(x, y, -1, h) ||
		f.isVisibleHorizontally(x, y, 1, h) || f.isVisibleHorizontally(x, y, -1, h)

}

func (f Forest) isVisibleVertically(x, y, dir, h int) bool {
	for i := y + dir; i >= 0 && i < len(f); i += dir {
		if f[i][x] >= h {
			return false
		}
	}
	return true
}

func (f Forest) isVisibleHorizontally(x, y, dir, h int) bool {
	for i := x + dir; i >= 0 && i < len(f[y]); i += dir {
		if f[y][i] >= h {
			return false
		}
	}
	return true
}

func (f Forest) visibilityVertically(x, y, dir, h int) int {
	distance := 0
	for i := y + dir; i >= 0 && i < len(f); i += dir {
		distance++
		if f[i][x] >= h {
			break
		}
	}
	return distance
}

func (f Forest) visibilityHorizontally(x, y, dir, h int) int {
	distance := 0
	for i := x + dir; i >= 0 && i < len(f[y]); i += dir {
		distance++
		if f[y][i] >= h {
			break
		}
	}
	return distance
}
