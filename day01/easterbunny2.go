package main

import (
	"fmt"
	"strconv"
	"strings"
)

type direction int

const (
	NORTH direction = 0 + iota
	EAST
	SOUTH
	WEST
)

type position struct {
	x         int
	y         int
	direction direction
}

type pos struct {
	x int
	y int
}

type move struct {
	rotation int
	distance int
}

var visited = map[pos]bool{}

func main() {
	input := "R3, L5, R2, L2, R1, L3, R1, R3, L4, R3, L1, L1, R1, L3, R2, L3, L2, R1, R1, L1, R4, L1, L4, R3, L2, L2, R1, L1, R5, R4, R2, L5, L2, R5, R5, L2, R3, R1, R1, L3, R1, L4, L4, L190, L5, L2, R4, L5, R4, R5, L4, R1, R2, L5, R50, L2, R1, R73, R1, L2, R191, R2, L4, R1, L5, L5, R5, L3, L5, L4, R4, R5, L4, R4, R4, R5, L2, L5, R3, L4, L4, L5, R2, R2, R2, R4, L3, R4, R5, L3, R5, L2, R3, L1, R2, R2, L3, L1, R5, L3, L5, R2, R4, R1, L1, L5, R3, R2, L3, L4, L5, L1, R3, L5, L2, R2, L3, L4, L1, R1, R4, R2, R2, R4, R2, R2, L3, L3, L4, R4, L4, L4, R1, L4, L4, R1, L2, R5, R2, R3, R3, L2, L5, R3, L3, R5, L2, R3, R2, L4, L3, L1, R2, L2, L3, L5, R3, L1, L3, L4, L3"

	//input := "R8, R4, R4, R8"

	commands := strings.Split(input, ",")
	var p position
	p.x = 0
	p.y = 0
	p.direction = NORTH

	moves := make([]move, len(commands))

	for i, s := range commands {
		moves[i] = getMove(s)
	}

	for _, move := range moves {
		alreadyVisited := p.move(move)
		if alreadyVisited {
			break
		}
	}
	fmt.Println(p)
	fmt.Println(abs(p.x) + abs(p.y))
}

func (p *position) move(m move) bool {

	d := (int(p.direction) + m.rotation) % 4

	if d == -1 {
		d = 3
	}

	p.direction = direction(d)

	for i := 0; i < m.distance; i++ {

		switch p.direction {
		case NORTH:
			p.y = p.y + 1
			break
		case SOUTH:
			p.y = p.y - 1
			break
		case EAST:
			p.x = p.x + 1
			break
		case WEST:
			p.x = p.x - 1
			break
		}

		position := pos{x: p.x, y: p.y}

		_, ok := visited[position]
		if ok {
			fmt.Println(ok)
			return true
		}
		visited[position] = true

	}
	return false

}

func getMove(input string) move {
	var m move = move{}
	s := strings.TrimSpace(input)
	d := s[0]
	if d == 'R' {
		m.rotation = 1
	} else {
		m.rotation = -1
	}
	l := s[1:len(s)]
	i, _ := strconv.ParseInt(l, 10, 32)
	m.distance = int(i)
	return m
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
