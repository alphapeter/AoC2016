package main

import (
	"fmt"
)



func main() {
	input := 3018458
	solve1(input)
	solve2_O_N2_(input)

	//solve2_O_NLogN_(5) // 2
	//solve2_O_NLogN_(6) // 3
	//solve2_O_NLogN_(7) // 5
	//solve2_O_NLogN_(8) // 7
	//solve2_O_NLogN_(9) // 9
	//solve2(10) // 1
	//solve2(11) // 2
	//solve2(100) // 19
	// WIP
	//solve2_O_NLogN_(3018458) //expected 1424135



}

// first part start

//efficient solution
func solve1b(count int){
	winner := 0
	for i:= 2; i <= count; i++{
		winner = (winner + 2) %i
	}
	fmt.Println(winner+1)
}

//solution using linked list
func solve1(count int){
	currentElf := initElfs(count)

	for currentElf.nextElf != currentElf {
		currentElf.packages = currentElf.packages + currentElf.nextElf.packages

		currentElf.nextElf = currentElf.nextElf.nextElf
		currentElf = currentElf.nextElf
	}

	currentElf.print()
}



type Elf struct {
	name     int
	packages int
	nextElf  *Elf
}

func (elf *Elf) print() {
	fmt.Printf("Elf %d, has %d packages\n", elf.name, elf.packages)
}

func initElfs(count int)*Elf{
	firstElf := &Elf{name: 1, packages: 1}
	currentElf := firstElf

	for i := 1; i < count; i++ {
		currentElf.nextElf = &Elf{name: i + 1, packages: 1}
		currentElf = currentElf.nextElf
	}
	currentElf.nextElf = firstElf
	return firstElf
}
// first part end

//second part O(N)
func solve2_O_N(count int){
	winner2 := 0
	for i:= 2; i<=count; i++{
		n_minus1_would_kill := i/2 - 1
		inc := 2
		if(n_minus1_would_kill > winner2){
			inc = 1
		}
		winner2 = (winner2 + inc) % i
	}
	fmt.Println(winner2+1)
}

// second part O(N^2) solution start using array
func solve2_O_N2_(count int){
	fmt.Println("-------Solving 19B, O(N^2)-------------")
	elements := make([]int, count);

	for i:=0; i<count; i++{
		elements[i] = i+1
	}
	t:= createList(count)
	i := 0;
	for t.Size > 1{
		opposite := (i + (t.Size)/2) % t.Size

		if( (i %t.Size) < opposite){
			i = (i + 1) % t.Size
		}

		t.remove(opposite)
	}
	fmt.Printf("part 2, last elf standing:%d\n", t.data[0])
}

func createList(size int) *List{
	list := List{Size:size, data:make([]int,size) }
	for i:= 0; i <size; i++ {
		list.data[i] = i+1
	}
	return &list
}

type List struct {
	Size int
	data []int
}

func (l *List) remove(element int){
	for i:= element; i < l.Size -1; i++{
		l.data[i] = l.data[i+1]
	}
	l.data[l.Size-1] = -1
	l.Size--
}

// END Solution 2 O(n^2)

// WIP, there's a bug in the rank tree
// second part O(N*LogN) (using rank tree)
func solve2_O_NLogN_(count int){
	fmt.Println("-------Solving 19B, O(NLog(N))-------------")
	elements := make([]int, count);

	for i:=0; i<count; i++{
		elements[i] = i+1
	}
	t := buildTree(elements,0, count-1)

	i := 0;
	for t.Size > 1{
		//fmt.Println("--")
		opposite := (i + (t.Size)/2) % t.Size
		//fmt.Println(t.find(i%t.Size))
		//fmt.Printf("index: %d, size: %d, oposite: %d |", i, t.Size, opposite)
		//fmt.Printf("i,o,s:{%d,%d, %d} %d -> %d\n", i, opposite, t.Size, t.find(i%t.Size).Value, t.find(opposite).Value)
		//fmt.Printf("i,o,s:{%d,%d, %d} %d -> %d\n", i, opposite, t.Size, t.data[i%t.Size], t.data[opposite])
		//fmt.Printf("i: %d, opposite:%d\n", i, opposite)
		//fmt.Printf("size: %d\n", t.Size)
		//fmt.Printf("i: %d, t.size/2: %d\n", i%t.Size, t.Size/2)

		//fmt.Println("----")

		if( (i %t.Size) < opposite){
			//fmt.Println("++")
			i = (i + 1) % t.Size
		}

		t = remove(t, opposite)
		//fmt.Println(t.data)
		//fmt.Println(i)
	}
	//fmt.Printf("part 2, last elf standing:%d\n", t.Value)
	fmt.Printf("part 2, last elf standing:%d\n", t.Value)
}

type Tree struct {
	Left *Tree
	Right *Tree
	Value int
	Size int
}

func (t *Tree) recomputeSize(){
	leftSize, rightSize := 0, 0
	if(t.Left != nil){
		leftSize = t.Left.Size
	}
	if(t.Right != nil){
		rightSize = t.Right.Size
	}
	t.Size = leftSize + 1 + rightSize
}

func buildTree(elements []int, start int, end int) *Tree {
	if(end < start){
		return nil;
	}

	if start == end {
		return &Tree{Value:elements[start], Size: 1}
	}
	mid := (start+end)/2

	t := Tree{Value:elements[mid]}
	t.Left = buildTree(elements, start, mid -1)
	t.Right = buildTree(elements, mid+1, end)
	t.recomputeSize()
	return &t;
}

func (t *Tree) find(index int) *Tree{
	leftSize := 0
	if(t.Left != nil){
		leftSize = t.Left.Size
	}
	if(index < leftSize){
		return t.Left.find(index)
	}
	if(index == leftSize){
		return t
	}
	return t.Right.find(index -leftSize - 1)
}

func remove(t *Tree, index int) *Tree{
	leftSize := 0
	if(t.Left != nil){
		leftSize = t.Left.Size
	}
	if(index < leftSize){
		t.Left = remove(t.Left, index)

	}else if(index > leftSize){
		t.Right = remove(t.Right, index - leftSize - 1)
	}else{
		if(t.Left != nil && t.Right != nil){
			t.Value = findMin(t.Right).Value
			t.Right = removeMin(t.Right)

		}else {
		if(t.Left != nil){
			t = t.Left
		}else{
			t = t.Right
		}
	}
	}
	if(t != nil){
		t.recomputeSize()
	}
	return t
}

func findMin (t *Tree) *Tree{
	if(t.Left == nil){
		return t
	}
	return findMin(t.Left)
}

func removeMin(t *Tree) *Tree{
	if(t.Left == nil){
		return t.Right
	}else {
		t.Left = removeMin(t.Left)
		return t
	}
}

//second part end




