package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Password []rune

func main() {
	inputData, _ := ioutil.ReadFile("day21/input.txt")
	input := strings.Split(string(inputData), "\n")
	solve1(input)
	solve2(input)
}

func solve1(input []string) {
	password := Password("abcdefgh")
	for _, row := range input {
		scamble(row, &password, false)

	}
	fmt.Printf("answer1: %s\n", string(password))
}

func solve2(input []string) {
	password := Password("fbgdceah")
	for i := len(input) - 1; i >= 0; i-- {
		row := input[i]
		scamble(row, &password, true)
	}
	fmt.Printf("answer2: %s\n", string(password))
}

func scamble(row string, password *Password, inverse bool) {
	if row == "" {
		return
	}

	command := strings.Split(row, " ")
	switch command[0] + "_" + command[1] {
	case "swap_letter":
		a := parseRune(command[2])
		b := parseRune(command[5])
		password.swapRune(a, b)
		break
	case "swap_position":
		i := parseInt(command[2])
		j := parseInt(command[5])
		password.swapIndex(i, j)
		break
	case "move_position":
		i := parseInt(command[2])
		j := parseInt(command[5])

		if inverse {
			password.move(j, i)
		} else {
			password.move(i, j)
		}
		break
	case "rotate_based":
		letter := parseRune(command[6])
		password.rotateBy(letter, inverse)
		break
	case "rotate_left":
		steps := parseInt(command[2])
		if inverse {
			password.rotateRight(steps)
		} else {
			password.rotateLeft(steps)
		}
		break
	case "rotate_right":
		steps := parseInt(command[2])
		if inverse {
			password.rotateLeft(steps)
		} else {
			password.rotateRight(steps)
		}
		break
	case "reverse_positions":
		i := parseInt(command[2])
		j := parseInt(command[4])
		password.reverse(i, j)
		break
	default:
		panic("not implemented")
	}
}

func (p *Password) swapIndex(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *Password) swapRune(a, b rune) {
	s := string(*p)
	i := strings.IndexRune(s, a)
	j := strings.IndexRune(s, b)
	p.swapIndex(i, j)
}

func (p *Password) rotateBy(r rune, inverse bool) {
	index := strings.IndexRune(string(*p), r)

	if inverse {
		var steps int = 0
		switch index {
		case 0:
			steps = 9
			break
		case 1:
			steps = 1
			break
		case 2:
			steps = 6
			break
		case 3:
			steps = 2
			break
		case 4:
			steps = 7
			break
		case 5:
			steps = 3
			break
		case 6:
			steps = 8
			break
		case 7:
			steps = 4
			break
		default:
			panic("case not considered")
		}
		p.rotateLeft(steps)

	} else {
		if index >= 4 {
			index++
		}
		steps := index + 1
		p.rotateRight(steps)
	}
}

func (p *Password) rotateLeftBy(r rune) {
	index := strings.IndexRune(string(*p), r)
	if index >= 4 {
		index++
	}
	steps := index + 1
	p.rotateRight(steps)
}

func (p *Password) rotateLeft(steps int) {
	steps = steps % len(*p)
	a := append((*p)[steps:], (*p)[0:steps]...)
	for i := 0; i < len(*p); i++ {
		(*p)[i] = a[i]
	}
}

func (p *Password) rotateRight(steps int) {
	steps = steps % len(*p)
	if steps == 0 {
		return
	}
	a := append((*p)[len(*p)-steps:], (*p)[0:len(*p)-steps]...)
	for i := 0; i < len(*p); i++ {
		(*p)[i] = a[i]
	}
}

func (p *Password) reverse(i, j int) {
	for i < j {
		p.swapIndex(i, j)
		j--
		i++
	}
}

func (p *Password) move(s, d int) {
	letter := p.remove(s)
	p.insert(d, letter)
}

func (p *Password) remove(index int) rune {
	letter := (*p)[index]
	for i := index; i < len(*p)-1; i++ {
		(*p)[i] = (*p)[i+1]
	}
	(*p)[len(*p)-1] = '#'
	return letter
}

func (p *Password) insert(index int, letter rune) {
	for i := len(*p) - 1; i > index; i-- {
		(*p)[i] = (*p)[i-1]
	}
	(*p)[index] = letter
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 64)
	return int(a)
}

func parseRune(s string) rune {
	return rune(s[0])
}
