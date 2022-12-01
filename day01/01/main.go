package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	elfs := strings.Split(input, "\n\n")

	var sums []int

	for _, elf := range elfs {
		items := strings.Split(elf, "\n")
		total := 0
		for _, item := range items {
			num, _ := strconv.Atoi(item)
			total += num
		}
		sums = append(sums, total)
	}

	sort.Ints(sums)
	log.Println(sums[len(sums)-1])

	top3 := sums[len(sums)-3:]
	top3Total := 0
	for _, i := range top3 {
		top3Total += i
	}
	log.Println(top3Total)
}
