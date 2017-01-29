package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day25/input.txt")
	input := string(inputData)
	rows := strings.Split(input, "\n")
	for i:= 0; i < 100000; i++{
		fmt.Printf("%d:",i)
		result := run(rows, map[string]int{"a": i})
		if result == "01010101010101010101"{
			fmt.Printf("answer1: %d",  i)
			return
		}
	}
}

func run(rows []string, registers map[string]int) string{
	outputs := 20
	result := []rune{}
	for pPointer := 0; pPointer < len(rows); pPointer++ {
		command := strings.Fields(rows[pPointer])
		if len(command) == 0 {
			continue
		}
		switch command[0] {
		case "cpy":
			source := command[1]
			destination := command[2]
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
			if val != 0 {
				pPointer = pPointer + parseInt(command[2]) - 1
			}
			break
		case "out":
			source := command[1]
			var val int
			if rune(source[0]) >= 'a' {
				val = registers[source]
			} else {
				val = parseInt(source)
			}
			result = append(result, rune(val+48))
			outputs--
			if outputs == 0 {
				fmt.Println(string(result))
				return string(result)
			}
			break

		default:
			panic(fmt.Sprintf("command not implemented %s", command))
		}
	}
	return ""
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
