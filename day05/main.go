package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {

	moves := strings.Split(moves, "\n")

	for _, m := range moves {
		NewMove(m).execute9000(stacks1)
		NewMove(m).execute9001(stacks2)
	}

	log.Println("----")

	topRow1 := ""
	for _, stack := range stacks1 {
		topRow1 += stack[0]
	}

	log.Println("Top row:", topRow1)
	log.Println("----")

	topRow2 := ""
	for _, stack := range stacks2 {
		topRow2 += stack[0]
	}

	log.Println("Top row:", topRow2)
}

type Stack []string

func NewMove(s string) Move {
	parts := strings.Split(s, " ")

	count, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])

	return Move{
		from:  from - 1,
		to:    to - 1,
		count: count,
	}
}

type Move struct {
	from  int
	to    int
	count int
}

func (m Move) execute9000(stacks []Stack) {
	for i := 0; i < m.count; i++ {
		tmp := stacks[m.from][0]
		stacks[m.from] = stacks[m.from][1:]
		stacks[m.to] = append([]string{tmp}, stacks[m.to]...)
	}
}

func (m Move) execute9001(stacks []Stack) {
	tmp := stacks[m.from][0:m.count]
	stacks[m.from] = stacks[m.from][m.count:]

	newTo := Stack{}
	newTo = append(newTo, tmp...)
	newTo = append(newTo, stacks[m.to]...)
	stacks[m.to] = newTo
}
