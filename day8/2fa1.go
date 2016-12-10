package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type display struct {
	rows    int
	columns int
	data    [][]byte
}

func main() {

	d := display{rows: 6, columns: 50}
	d.init()

	inputData, _ := ioutil.ReadFile("day8/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")
	fmt.Println("number of commands: ", len(rows))

	//d.rect(3,2)
	//d.print()
	//fmt.Println("-")
	//d.rotateColumn(1,1)
	//d.print()
	//fmt.Println("-")
	//d.rotateRow(0,4)
	//d.print()
	//fmt.Println("-")
	//
	//d.rotateColumn(1,1)

	for _, row := range rows {
		if row == "" {
			continue
		}

		d.executeCommand(row)
	}
	d.print()
	fmt.Println("answer: ", d.countPixels())
}

func (d *display) init() {
	d.data = make([][]byte, d.rows)
	for i := 0; i < d.rows; i++ {
		d.data[i] = make([]byte, d.columns)
	}
	d.reset()
}

func (d *display) reset() {
	for i := 0; i < len(d.data); i++ {
		for j := 0; j < len(d.data[i]); j++ {
			d.data[i][j] = '.'
		}
	}
}
func (d *display) executeCommand(s string) {
	fields := strings.Fields(s)
	switch fields[0] {
	case "rect":
		size := strings.Split(fields[1], "x")
		d.rect(parseInt(size[0]), parseInt(size[1]))
		break
	case "rotate":
		if strings.Contains(fields[1], "row") {
			//rotate row y=4 by 20
			row := parseInt(fields[2][2:])
			pixelCount := parseInt(fields[4])
			d.rotateRow(row, pixelCount)
		} else {
			//rotate column x=13 by 3
			column := parseInt(fields[2][2:])
			pixelCount := parseInt(fields[4])
			d.rotateColumn(column, pixelCount)
		}
		break
	}
}

func (d *display) rect(cols int, rows int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			d.data[i][j] = '#'
		}
	}
}

func (d *display) rotateColumn(colIndex int, pixelCount int) {
	column := make([]byte, len(d.data))
	for i := 0; i < d.rows; i++ {
		column[i] = d.data[i][colIndex]
	}

	newcolumn := append(column[d.rows-pixelCount:], column[0:d.rows-pixelCount]...)

	for i := 0; i < d.rows; i++ {
		d.data[i][colIndex] = newcolumn[i]
	}

}
func (d *display) rotateRow(row int, pixelCount int) {
	d.data[row] = append(d.data[row][d.columns-pixelCount:], d.data[row][0:d.columns-pixelCount]...)
}

func (d *display) print() {
	for i := 0; i < d.rows; i++ {
		fmt.Printf("%s\n", d.data[i])
	}

}

func (d *display) countPixels() int {
	sum := 0
	for i := 0; i < len(d.data); i++ {
		for j := 0; j < len(d.data[i]); j++ {
			if d.data[i][j] == '#' {
				sum++
			}
		}
	}
	return sum
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
