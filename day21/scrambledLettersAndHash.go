package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Password []rune

func main() {
	inputData, _ := ioutil.ReadFile("day21/input.txt")
	input := strings.Split(string(inputData), "\n")
	//WIP
	for _, row := range input {
		if row == "" {
			continue
		}
		//swap index - swap position 2 with position 7
		//swap letters - swap letter a with letter d

		//rotate - rotate based on position of letter h
		//rotate - rotate left 2 steps
		//rotate - rotate right 2 steps

		//move - move position 6 to position 3

		//reverse - reverse positions 4 through 6
	}
}

func (p *Password) swapIndex(i, j int){
	p[i], p[j] = p[j], p[i]
}

func (p *Password) swapRune(a, b rune){
	s := string(*p)
	i := strings.IndexRune(s, a)
	j := strings.IndexRune(s, b)
	p.swapIndex(i,j)
}

func (p* Password) rotateBy(r rune){
	index := strings.IndexRune(string(p), r) + 1
	if index >= 4{
		index++
	}
	steps := index + 1
	p.rotateRight(steps)
}

//TODO: fix
func (p* Password) rotateLeft(steps int){
	a := &append(p[len(*p)-steps:], p[0:len(*p)-steps]...)
	for i:=0; i<len(*p); i++{
		p[i] = a[i]
	}
}
func (p* Password) rotateRight(steps int){
	a := &append(p[len(*p)-steps:], p[0:len(*p)-steps]...)
	for i:=0; i<len(*p); i++{
		p[i] = a[i]
	}
}
func(p* Password) reverse(i, j int){
	for i < j {
		p.swapIndex(i,j)
		j--
		i++
	}
}

//TODO: fix
func(p* Password) move(s, d int){
	tmp := p[s]
	for i := s; i < len(p){
		p[i] = p[i+1]
	}
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 64)
	return int(a)
}
