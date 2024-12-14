package main

import (
	"fmt"
	"strconv"
)

func NewGrid(data []string) *Grid {
	g := &Grid{}
	g.data = []Node{}
	g.w = len(data[0])
	g.h = len(data)
	g.vidx = 1
	for y,row := range data {
		for x,v := range row {
			value, err := strconv.Atoi(string(v))
			check(err)
			g.data = append(g.data, Node{x: x,y: y,value: value,visited: g.vidx-1})
		}
	}
	return g 
}

func (g *Grid) WithinBounds(x,y int) bool {
	return x > -1 && x < g.w && y > -1 && y < g.h
}
func (g *Grid) Get(x, y int) *Node {
	return &g.data[x + y * g.h]
}
func (g *Grid) GetNeighbours(x, y int) []*Node {
	offsets := []int {0,1,0,-1,1,0,-1,0}
	neighbours := []*Node{}
	for o := 0; o < len(offsets); o+=2 {
		nx := x + offsets[o]; ny := y + offsets[o+1]
		if g.WithinBounds(nx,ny) {neighbours = append(neighbours, g.Get(nx, ny))}
	}
	return neighbours
}
type Node struct {
	x, y, value, visited int
}
type Grid struct {
	w,h, vidx int
	data []Node
}
func (g *Grid) DFS(v *Node,isPartTwo bool) int{
	s := []*Node{}
	s = append(s, v)
	count := 0
	for len(s) != 0 {
		v = s[len(s)-1] 
		s = s[:len(s)-1] // popping last element
		if g.vidx == v.visited && !isPartTwo {continue} // continue if its already visited
		v.visited = g.vidx // label v as visited 
		if v.value == 9 {count++;continue}
		neighbours := g.GetNeighbours(v.x, v.y)
		for _,n := range neighbours {
			if v.value + 1 == n.value {s = append(s, n)}
		}
	}
	g.vidx++
	return count
}
func main(){
	defer timer("program")()
	// Input_Parsing
	data := []string{}
	for v := range readLine() {data = append(data, v)}
	grid := NewGrid(data)
	// Part_One		
	sum := 0
	for _,v := range grid.data {
		if v.value == 0 {sum += grid.DFS(&v,false)}
	}		
	fmt.Println("Part_One _>",sum)
	// Part_Two
	sum = 0
	for _,v := range grid.data {
		if v.value == 0 {sum += grid.DFS(&v,true)}
	}
	fmt.Println("Part_Two _>",sum)
}
