package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")

	totalContained := 0
	totalOverlapped := 0
	for _, line := range lines {
		e1, e2 := split(line)
		r1 := asRange(e1)
		r2 := asRange(e2)

		if r1.contains(r2) || r2.contains(r1) {
			totalContained++
		}

		if r1.overlaps(r2) || r2.overlaps(r1) {
			totalOverlapped++
		}
	}
	log.Println("Totally contained:", totalContained)
	log.Println("Totally overlapped:", totalOverlapped)
}

func split(s string) (string, string) {
	parts := strings.SplitN(s, ",", 2)
	return parts[0], parts[1]
}

func asRange(s string) Range {
	parts := strings.Split(s, "-")
	from, _ := strconv.Atoi(parts[0])
	to, _ := strconv.Atoi(parts[1])

	return Range{
		from: from,
		to:   to,
	}
}

type Range struct {
	from int
	to   int
}

func (r Range) contains(other Range) bool {
	return r.from <= other.from && r.to >= other.to
}

func (r Range) overlaps(other Range) bool {
	return (r.from <= other.from && r.to >= other.from) || (r.from <= other.to && r.to >= other.to)
}
