package main

import (
	"log"
	"strings"
)

func main() {

	p1()
	p2()

}

func p2() {
	rucksacks := strings.Split(input, "\n")

	total := 0
	for i := 0; i < len(rucksacks); i += 3 {
		elf1 := rucksacks[i]
		elf2 := rucksacks[i+1]
		elf3 := rucksacks[i+2]

		whatTheyAllHave := doubleUnion(elf1, elf2, elf3)
		total += prio(whatTheyAllHave)
	}

	log.Println(total)
}

func p1() {
	rucksacks := strings.Split(input, "\n")

	p := 0
	for _, rucksack := range rucksacks {
		a, b := split(rucksack)
		both := union(a, b)
		p += prio(both)
	}
	log.Println("Prio total: ", p)
}

func union(a, b string) string {
	for _, l := range a {
		if strings.ContainsRune(b, l) {
			return string(l)
		}
	}
	return ""
}

func doubleUnion(a, b, c string) string {
	for _, l := range a {
		if strings.ContainsRune(b, l) && strings.ContainsRune(c, l) {
			return string(l)
		}
	}
	return ""
}

func split(in string) (string, string) {
	return in[0 : len(in)/2], in[len(in)/2 : len(in)]
}

func prio(s string) int {
	var priorities = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
	return priorities[s]
}
