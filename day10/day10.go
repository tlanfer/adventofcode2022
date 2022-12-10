package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func NewComputer() *Computa {
	return &Computa{
		x: 1,
	}
}

type Computa struct {
	cycles int
	x      int

	signalSum int
}

func (d *Computa) Run(input string) int {
	instructions := strings.Split(input, "\n")

	for _, instruction := range instructions {
		ps := strings.Split(instruction, " ")

		cmd := ps[0]

		switch cmd {
		case "noop":
			d.noop()
		case "addx":
			d.addx(ps[1])
		}
	}

	fmt.Println("")

	return d.signalSum
}

func (d *Computa) clock() {
	d.pixel()
	d.cycles++
	if (d.cycles-20)%40 == 0 {
		signalStrength := d.cycles * d.x
		d.signalSum += signalStrength
		//log.Printf("Signal strength at %v * %v = %v (Total %v)", d.cycles, d.x, signalStrength, d.signalSum)
	}
}

func (d *Computa) noop() {
	d.clock()
}

func (d *Computa) addx(s string) {
	d.clock()
	d.clock()
	x, _ := strconv.Atoi(s)
	d.x += x
}

func (d *Computa) pixel() {

	pos := d.cycles % 40

	if pos >= d.x-1 && pos <= d.x+1 {
		fmt.Print("█")
	} else {
		fmt.Print(" ")
	}

	if pos == 39 {
		fmt.Print("\n")
	}
}
