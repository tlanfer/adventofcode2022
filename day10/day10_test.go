package day10_test

import (
	"adventofcode2022/day10"
	"testing"
)

func TestDay10(t *testing.T) {

	inputs := map[string]struct {
		input    string
		expected int
	}{
		"test":    {input: testInput, expected: 13140},
		"actual":  {input: actualInput, expected: 15220},
		"fireeye": {input: fireeyeinput, expected: 12740},
	}

	for name, tc := range inputs {
		t.Run(name, func(t *testing.T) {
			d := day10.NewComputer()
			actual := d.Run(tc.input)

			if actual != tc.expected {
				t.Fatalf("signal strength expected = %v, was %v", tc.expected, actual)
			}
		})
	}

}

const testInput = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

const actualInput = `noop
noop
noop
addx 5
addx 1
addx 4
addx 1
noop
addx 4
noop
addx 1
addx 4
addx 8
addx -7
addx 3
addx 1
noop
addx 4
addx 2
addx 5
addx -1
noop
addx -37
noop
noop
addx 3
addx 2
addx 13
addx 12
addx -15
addx -2
addx 2
addx -11
addx 18
addx 2
addx -15
addx 16
addx 5
addx 2
addx 5
noop
noop
noop
addx 3
addx -2
addx -38
noop
addx 3
addx 4
noop
noop
noop
noop
noop
addx 5
addx 5
noop
noop
addx 21
addx -17
addx 6
noop
noop
noop
noop
addx 5
noop
noop
noop
noop
noop
addx 3
addx 5
addx -38
noop
noop
addx 5
addx -2
addx 1
addx 7
noop
addx 22
addx -18
addx -11
addx 27
addx -13
addx 2
addx 5
addx -8
addx 9
addx 2
noop
addx 7
noop
addx 1
noop
addx -38
noop
addx 2
addx 5
addx -3
noop
addx 8
addx 11
addx -6
noop
addx 24
addx -31
addx 10
addx 2
addx 5
addx 3
noop
addx 2
addx -29
addx 21
addx 11
addx 5
addx -39
addx 4
addx -2
addx 2
addx 7
noop
addx -1
addx 2
noop
addx 4
noop
addx 1
addx 2
addx 5
addx 2
noop
noop
addx -6
addx 9
addx -18
addx 25
addx 3
noop
addx -17
noop`

const fireeyeinput = `noop
noop
noop
addx 4
addx 1
addx 5
addx 1
addx 5
noop
addx -1
addx -6
addx 11
noop
noop
noop
noop
addx 6
addx 5
noop
noop
noop
addx -30
addx 34
addx 2
addx -39
noop
addx 5
addx 2
addx 19
addx -18
addx 2
addx 5
addx 2
addx 3
noop
addx 2
addx 3
noop
addx 2
addx 3
noop
addx 2
addx 3
noop
addx 2
addx -15
addx -22
noop
noop
addx 5
addx 2
noop
noop
addx 14
addx -11
addx 5
addx 2
addx 3
noop
addx 2
addx -16
addx 17
addx 2
addx 5
addx 2
addx -6
addx -25
addx 35
addx 1
addx -36
addx 1
addx 22
addx -19
addx 5
addx 2
noop
noop
addx 5
noop
noop
noop
addx 1
addx 4
noop
noop
noop
addx 5
noop
addx 1
addx 2
addx 3
addx 4
addx -34
addx 21
addx -24
addx 2
addx 5
addx 7
addx -6
addx 2
addx 30
addx -23
addx 10
addx -9
addx 2
addx 2
addx 5
addx -12
addx 13
addx 2
addx 5
addx 2
addx -12
addx -24
addx -1
noop
addx 3
addx 3
addx 1
addx 5
addx 21
addx -16
noop
addx 19
addx -18
addx 2
addx 5
addx 2
addx 3
noop
addx 3
addx -1
addx 1
addx 2
addx -18
addx 1
noop`