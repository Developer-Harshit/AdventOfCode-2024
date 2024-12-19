package main

import (
	"fmt"
	"regexp"
	"strings"
)

func GetStringFromRobotMap(r *[]int, w, h int) string {
	m := map[int]bool{}
	for i := 0; i < len(*r); i+=4 {
		idx := (*r)[i] + (*r)[i+1] * w
		_, ok := m[idx]
		if !ok {m[idx] = true}
	}
	str := ""
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			_, ok := m[i + j*w]
			if !ok {str += " "} else {str += "#"}
		}
		str += "\n"
	}
	return str

}
func IterateRobots(r *[]int,w, h, dt int) {
	for i:=0;i<len(*r);i+=4 {
		x  := (*r)[i+0];      y  := (*r)[i+1]
		x += (*r)[i+2] * dt ; y += (*r)[i+3] * dt
		x = x % w; y = y % h
		if x < 0 {x += w} else if x >= w {x -= w}
		if y < 0 {y += h} else if y >= h {y -= h}
		(*r)[i+0] = x; (*r)[i+1] = y
	}
}
func main(){
	defer timer("program")()
	// Input_Parsing
	width := 101; height := 103
	data := []int{}
	re := regexp.MustCompile("-?[0-9]+")
	for v := range readLine() {
		vals := sliceAtoi(re.FindAllString(v,-1)) 
		data = append(data[:], vals...)
	}
	// Part_One
	kw := width/2; kh := height/2
	quadrants := []int{
		0,0,kw+1,0,0,kh+1,kw+1,kh+1,
	}
	IterateRobots(&data,width,height,100)
	qScore := [4]int{0,0,0,0}
	for k := 0; k < len(data); k+=4 {
		x := data[k+0]; y := data[k+1]	
		for q := 0; q < len(quadrants); q+=2 {
			w := quadrants[q+0]; h := quadrants[q+1]
			if x < w || y < h || x >= w+kw || y >= h+kh {continue} 
			qScore[q/2]++
		}
	}
	fmt.Println("Part_One _>",qScore[0]*qScore[1]*qScore[2]*qScore[3])
	// Part_Two
	s := 100 
	for ; ; s++ {
		IterateRobots(&data, width, height, 1)		
		str := GetStringFromRobotMap(&data, width, height)
		if strings.Contains(str, "#########") {
			fmt.Print("\033[H\033[2J","Iteration _>",s,"\n",str)
			break;
		}
	}
	fmt.Println("Part_Two _>",s+1)
}
