package day12

import (
	"fmt"
	"strings"
)

type Map struct {
	Elevation [][]int
	Start     Position
	End       Position
}

type Position struct {
	X, Y int
}

func Parse(input string) *Map {
	m := &Map{}
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		var row []int
		for x, r := range line {
			switch r {
			case 'S':
				m.Start = Position{x, y}
				row = append(row, 0)
			case 'E':
				m.End = Position{x, y}
				row = append(row, 25)
			default:
				e := int(r - 'a')
				row = append(row, e)
			}
		}
		m.Elevation = append(m.Elevation, row)
	}
	return m
}

func (m *Map) Hike() int {
	lowestDistance := 100000000
	for y, row := range m.Elevation {
		for x, loc := range row {
			if loc == 0 {
				dist := m.nav(Position{x, y})
				if dist > 0 && dist < lowestDistance {
					lowestDistance = dist
				}
			}
		}
	}
	return lowestDistance
}

func (m *Map) Navigate() int {
	return m.nav(m.Start)
}

func (m *Map) nav(startPoint Position) int {

	openSet := prio{}
	openSet.put(startPoint, distance(startPoint, m.End))

	pathTo := map[Position]Position{}
	scores := map[Position]int{
		startPoint: 0,
	}

	for {

		currentEl := openSet.take()
		if currentEl == nil {
			return -1
		}

		current := *currentEl

		if current == m.End {
			fullPath := walkBack(current, pathTo)
			//log.Println("Path:", len(fullPath))
			//log.Println("Path:", fullPath)
			//render(m.Elevation, fullPath)
			return len(fullPath)
		}

		neighbors := m.neighbors(current)

		for _, n := range neighbors {

			if m.Elevation[n.Y][n.X] > m.Elevation[current.Y][current.X]+1 {
				continue
			}

			tentativeScore := scores[current] + distance(current, n)
			if neighborScore, exists := scores[n]; !exists || tentativeScore < neighborScore {
				pathTo[n] = current
				scores[n] = tentativeScore

				if !openSet.contains(n) {
					openSet.put(n, tentativeScore+distance(n, m.End))
				}
			}
		}

	}
}

func walkBack(current Position, pathTo map[Position]Position) []Position {
	fullPath := []Position{current}
	for {
		prev, exists := pathTo[fullPath[0]]
		if !exists {
			return fullPath[1:]
		}

		fullPath = append([]Position{prev}, fullPath...)
	}
}

func render(elevation [][]int, path []Position) {
	for _, position := range path {
		elevation[position.Y][position.X] = -1
	}

	for _, line := range elevation {
		for _, c := range line {
			if c == -1 {
				fmt.Print("█")
			} else {
				fmt.Print(string(rune('a' + c)))
			}
		}
		fmt.Println()
	}
}

func distance(a, b Position) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func (m *Map) isValidDirection(from, to Position) bool {
	return m.Elevation[from.Y][from.X] >= (m.Elevation[to.Y][to.X] - 1)
}

func (m *Map) contains(path []Position, position Position) bool {
	for _, p := range path {
		if p == position {
			return true
		}
	}
	return false
}

func (m *Map) neighbors(p Position) []Position {
	var neighbors []Position

	if p.X > 0 {
		neighbors = append(neighbors, Position{p.X - 1, p.Y})
	}

	if p.Y > 0 {
		neighbors = append(neighbors, Position{p.X, p.Y - 1})
	}

	if p.X < len(m.Elevation[0])-1 {
		neighbors = append(neighbors, Position{p.X + 1, p.Y})
	}

	if p.Y < len(m.Elevation)-1 {
		neighbors = append(neighbors, Position{p.X, p.Y + 1})
	}

	return neighbors
}

type prio struct {
	items map[Position]int
}

func (p *prio) put(position Position, prio int) {
	if p.items == nil {
		p.items = map[Position]int{}
	}
	p.items[position] = prio
}

func (p *prio) take() *Position {
	if len(p.items) == 0 {
		return nil
	}

	lowestPrio := -1
	var lowestPosition *Position

	for position, prio := range p.items {
		if lowestPrio == -1 || prio <= lowestPrio {
			lowestPosition = &position
			lowestPrio = prio
		}
	}

	delete(p.items, *lowestPosition)

	return lowestPosition
}

func (p *prio) contains(position Position) bool {
	_, exists := p.items[position]
	return exists
}
