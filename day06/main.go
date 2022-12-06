package main

import "log"

func main() {
	pos := detect(input, 14)
	log.Println("Detected at", pos)
}

func detect(s string, seq int) int {
	i := seq

	for ; i < len(s); i++ {
		letters := map[rune]any{}
		for _, l := range s[i-seq : i] {
			letters[l] = struct{}{}
		}
		if len(letters) == seq {
			return i
		}
	}
	return -1
}
