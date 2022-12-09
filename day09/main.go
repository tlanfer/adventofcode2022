package main

import (
	"fmt"
	"log"
	"math"
)

const (
	startX = 12
	startY = 16
)

func main() {
	steps := Parse(input)

	s := snek{
		knots: []position{
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
			position{startX, startY},
		},
		tailPositions: map[string]any{},
	}

	for _, step := range steps {
		s.move(step.direction, step.distance)
		//s.print()
	}

	log.Printf("Tail visited %v positions", len(s.tailPositions))
}

type snek struct {
	knots         []position
	tailPositions map[string]any
}

func (s *snek) move(dir string, steps int) {
	log.Printf("Move %v to %v", steps, dir)
	tail := &s.knots[len(s.knots)-1]
	s.tailPositions[fmt.Sprintf("%04d-%04d", tail.x, tail.y)] = ""
	for i := 0; i < steps; i++ {
		// Move head
		switch dir {
		case "L":
			s.knots[0].x -= 1
		case "R":
			s.knots[0].x += 1
		case "U":
			s.knots[0].y -= 1
		case "D":
			s.knots[0].y += 1
		}

		// Move tails
		for k := 1; k < len(s.knots); k++ {
			prevKnot := &s.knots[k-1]
			knot := &s.knots[k]

			if prevKnot.distance(knot) > 1 {
				// move tail

				if prevKnot.x > knot.x {
					knot.x += 1
				} else if prevKnot.x < knot.x {
					knot.x -= 1
				}

				if prevKnot.y > knot.y {
					knot.y += 1
				} else if prevKnot.y < knot.y {
					knot.y -= 1
				}
			}
		}
		s.tailPositions[fmt.Sprintf("%04d-%04d", tail.x, tail.y)] = ""
	}
}

func (s *snek) print() {
	fmt.Printf("==== H (%v,%v)\n", s.knots[0].x, s.knots[0].y)
	for y := 0; y < 20; y++ {
		for x := 0; x < 27; x++ {
			mark := "."
			for i, k := range s.knots {
				if k.x == x && k.y == y {
					mark = fmt.Sprint(i)
				}
			}
			fmt.Print(mark)
		}
		fmt.Print("\n")
	}
	fmt.Print("====\n\n")
}

type position struct {
	x, y int
}

func (p position) distance(p2 *position) int {
	distanceX := math.Abs(float64(p.x - p2.x))
	distanceY := math.Abs(float64(p.y - p2.y))
	return int(math.Max(distanceX, distanceY))

}
