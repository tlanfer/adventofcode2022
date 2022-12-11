package day11

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func NewZoo() *Zoo {
	return &Zoo{
		Monkeys:        []Monkey{},
		CommonMultiple: 1,
	}
}

type Zoo struct {
	Monkeys        []Monkey
	CommonMultiple int64
}

func (z *Zoo) Parse(input string) {
	monkeyBlocks := strings.Split(input, "\n\n")
	for _, monkey := range monkeyBlocks {
		z.AddMonkey(monkey)
	}
}

func (z *Zoo) AddMonkey(m string) {
	lines := strings.Split(m, "\n")

	monkey := Monkey{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "Monkey "):
			//monkeyNo = parseint64(strings.TrimSuffix(strings.TrimPrefix(line, "Monkey "), ":"))

		case strings.HasPrefix(line, "Starting items: "):
			itemList := strings.TrimPrefix(line, "Starting items: ")
			itemSlice := strings.Split(itemList, ", ")
			for _, s := range itemSlice {
				monkey.Items = append(monkey.Items, parseInt(s))
			}

		case strings.HasPrefix(line, "Operation: new = old"):
			operationStr := strings.TrimPrefix(line, "Operation: new = old ")
			switch operationStr[0] {
			case '+':
				monkey.Worry = func(i int64) int64 {
					operand := i
					if operationStr[2:] != "old" {
						operand = parseInt(operationStr[2:])
					}
					return i + operand
				}
			case '*':
				monkey.Worry = func(i int64) int64 {
					//log.Print64f("%v * %v", i, parseint64(operationStr[2:]))
					operand := i
					if operationStr[2:] != "old" {
						operand = parseInt(operationStr[2:])
					}
					return i * operand
				}
			}

		case strings.HasPrefix(line, "Test: divisible by "):
			divisor := parseInt(strings.TrimPrefix(line, "Test: divisible by "))
			z.CommonMultiple *= divisor
			monkey.Test = func(i int64) bool {
				return i%divisor == 0
			}

		case strings.HasPrefix(line, "If true: throw to monkey "):
			monkey.TrueMonkey = parseInt(strings.TrimPrefix(line, "If true: throw to monkey "))

		case strings.HasPrefix(line, "If false: throw to monkey "):
			monkey.FalseMonkey = parseInt(strings.TrimPrefix(line, "If false: throw to monkey "))

		}
	}

	z.Monkeys = append(z.Monkeys, monkey)
}

func (z *Zoo) MonkeyBusiness(relaxFactor int64, rounds int) int64 {

	for i := 0; i < rounds; i++ {

		for mNo, monkey := range z.Monkeys {

			for _, item := range monkey.Items {

				newWorry := monkey.Worry(item)
				newWorry /= relaxFactor
				newWorry = newWorry % z.CommonMultiple

				newMonkey := monkey.FalseMonkey
				if monkey.Test(newWorry) {
					newMonkey = monkey.TrueMonkey
				}

				z.Monkeys[newMonkey].Items = append(z.Monkeys[newMonkey].Items, newWorry)
				z.Monkeys[mNo].Inspections++
			}
			z.Monkeys[mNo].Items = []int64{}
		}

		//log.Print64f("=== Round %v", i)
		//for n, monkey := range z.Monkeys {
		//	log.Print64f("Monkey %v: %v", n, monkey.Items)
		//}
	}

	sort.SliceStable(z.Monkeys, func(i, j int) bool {
		return z.Monkeys[i].Inspections > z.Monkeys[j].Inspections
	})

	for n, monkey := range z.Monkeys {
		log.Printf("Monkey %v inspected items %v times.", n, monkey.Inspections)
	}

	return z.Monkeys[0].Inspections * z.Monkeys[1].Inspections
}

func parseInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

type Monkey struct {
	Items       []int64
	Worry       func(int64) int64
	Test        func(int64) bool
	TrueMonkey  int64
	FalseMonkey int64
	Inspections int64
}
