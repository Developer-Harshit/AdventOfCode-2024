package main 
import (
	"fmt"
)
type Plot struct {
	x, y int
	visited bool
	value byte
}
func (p *Plot) floodFill(data [][]*Plot) (int, int) {
	if p.visited {return 0,0}
	openSet := []*Plot{p}
	w := len(data[0]); 
	h := len(data)
	offset := []int{0,1,1,0,0,-1,-1,0}
	dOffset := []int{-1,-1,-1,1,1,-1,1,1}
	area := 0
	perimeter := 0
	sides := 0
	for len(openSet) > 0 {
		v := openSet[0]
		openSet = openSet[1:] // pop first element
		if v.value != p.value { // if not of same kind
			perimeter++
			continue
		}
		if v.visited {continue}
		v.visited = true
		area++
		// for part two 
		for o := 0; o < len(dOffset); o+=2 {
			ox := dOffset[o]; oy := dOffset[o+1]
			dx := ox + v.x; dy := oy + v.y 
			hx := ox + v.x; hy := 0  + v.y 
			vx := 0  + v.x; vy := oy + v.y 
			dNotSame := dx > -1 && dy > -1 && dx < w && dy < h && data[dy][dx].value != v.value
			hSame    := hx > -1 && hy > -1 && hx < w && hy < h && data[hy][hx].value == v.value
			vSame    := vx > -1 && vy > -1 && vx < w && vy < h && data[vy][vx].value == v.value
			if (hSame && vSame && dNotSame) || (!hSame && !vSame) {
				sides++
			}
		}
		// for flood fill 
		for o := 0; o < len(offset); o+=2 {
			nx := offset[o] + v.x; ny := offset[o+1] + v.y
			if nx > -1 && ny > -1 && nx < w && ny < h { // if within bounds
				openSet = append(openSet, data[ny][nx])
			} else {
				perimeter++
			}
		}
	}
	return area * perimeter, area * sides 	
}

func main(){
	defer timer("program")()
	// Input_Parsing
	data := [][]*Plot{}
	j := 0
	for row := range readLine() {
		data = append(data, []*Plot{})
		for i,v := range row {
			data[j] = append(data[j], &Plot{value: byte(v),x:i,y:j,visited:false})
		}
		j++
	}
	// Part_One & Part_Two
	sum1 := 0
	sum2 := 0
	for _,row := range data {
		for _,v := range row {
			one,two := v.floodFill(data)
			sum1 += one
			sum2 += two
		}
	}
	fmt.Println("Part_One _>",sum1)
	fmt.Println("Part_Two _>",sum2)
}
