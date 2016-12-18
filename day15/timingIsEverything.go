package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type State []*Disc

type Disc struct {
	positions int
	position  int
}

func (s *Disc) tick() {
	s.position = (s.position + 1) % s.positions
}

func (s *Disc) isValid() bool {
	return s.position == 0
}

func (s *State) tick() {
	for _, disc := range *s {
		disc.tick()
	}
}

func (s *State) print() {
	fmt.Println("-------------------")
	for _, disc := range *s {
		fmt.Printf("%d/%d\n", disc.position, disc.positions)
	}
	fmt.Println("-------------------")
}

func (s *State) isValid() bool {
	for _, disc := range *s {
		if !disc.isValid() {
			return false
		}
	}
	return true
}

func main() {
	inputData, _ := ioutil.ReadFile("day15/input.txt")
	input := strings.Split(strings.TrimSpace(string(inputData)), "\n")
	solve(input, 1)
	solve(append(input, "Disc #7 has 11 positions; at time=0, it is at position 0."), 2)

}
func solve(input []string, problemNo int) {
	positionsRegex := regexp.MustCompile(`(\d+) positions`)
	positionRegex := regexp.MustCompile(`position (\d+).`)
	state := State(make([]*Disc, len(input)))

	for i, row := range input {
		positions := parseInt(positionsRegex.FindStringSubmatch(row)[1])
		position := parseInt(positionRegex.FindStringSubmatch(row)[1])
		state[i] = &Disc{position: (position + i) % positions, positions: positions}
	}
	for i := 0; true; i++ {
		state.tick()

		if state.isValid() {
			state.print()
			fmt.Printf("Answer%d: %d ticks\n", problemNo, i)
			return
		}
	}
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
