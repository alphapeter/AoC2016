package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day23/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")
	fmt.Printf("answer 1: %d\n", run(rows, map[string]int{"a": 7}))
	//fmt.Printf("answer 2: %d\n", run(rows, map[string]int{"c": 1}))
}

func run(rows []string, registers map[string]int) int {

	for pPointer := 0; pPointer < len(rows); pPointer++ {
		command := strings.Fields(rows[pPointer])
		if len(command) == 0 {
			continue
		}
		switch command[0] {
		case "cpy":
			source := command[1]
			destination := command[2]
			if rune(destination[0]) <= '9' || rune(destination[0]) == '-' {
				continue
			}
			var val int
			if rune(source[0]) >= 'a' {
				val = registers[source]
			} else {
				val = parseInt(source)
			}
			registers[destination] = val
			break

		case "inc":
			registers[command[1]] = registers[command[1]] + 1
			break
		case "dec":
			registers[command[1]] = registers[command[1]] - 1
			break
		case "jnz":
			var val int
			if rune(command[1][0]) >= 'a' {
				val = registers[command[1]]
			} else {
				val = parseInt(command[1])
			}
			var offset int
			if rune(command[2][0]) >= 'a' {
				offset = registers[command[2]]
			} else {
				offset = parseInt(command[2])
			}
			if val != 0 {
				pPointer = pPointer + offset - 1
			}
			break
		case "tgl":
			insPointer := pPointer + registers[command[1]]
			if insPointer >= len(rows) || insPointer <= 0 {
				continue
			}

			tglRow := strings.Fields(rows[insPointer])
			tglCommand := tglRow[0]
			tglArguments := tglRow[1:]

			if len(tglArguments) == 1 {
				if tglCommand == "inc" {
					rows[insPointer] = "dec " + tglArguments[0]
				} else {
					rows[insPointer] = "inc " + tglArguments[0]
				}
			} else { //binary
				if tglCommand == "jnz" {
					rows[insPointer] = "cpy " + tglArguments[0] + " " + tglArguments[1]
				} else {
					rows[insPointer] = "jnz " + tglArguments[0] + " " + tglArguments[1]
				}
			}
		default:
			panic(fmt.Sprintf("command not implemented %s", command))
		}
	}
	return registers["a"]
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
