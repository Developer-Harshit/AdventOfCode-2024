package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)
func main(){
	defer timer("program")();
	// Input_Parsing
	leftList := []int{};
	rightList := []int{};
	for v := range readLine() {

		f := strings.Fields(v);
		if(len(f) != 2)	{

			log.Fatal("Invalid input");
		}
		a,err := strconv.Atoi(f[0]);
		check(err);
		b,err := strconv.Atoi(f[1]);
		check(err);
		leftList = append(leftList, a);
		rightList = append(rightList, b);
	}
	// Part_One 
	sort.Ints(leftList);
	sort.Ints(rightList);
	sum := 0;
	for i := 0; i < len(leftList); i++ {

		val := leftList[i] - rightList[i];
		if (val < 0) {
			val*= -1;
		}
		sum += val;
	}
	fmt.Println("Part_One _>",sum);
	// Part_Two
	score := 0
	for _,v1 := range leftList {
	
		for _,v2 := range rightList {
			
			if v1 == v2 {
				score += v1;
			}
		}
	}
	fmt.Println("Part_Two _>",score);
}
