package main

import (
	"fmt"
	"strings"
)
func main(){
	defer timer("program")();
	// Input_Parsing
	data := strings.Split(readRaw(), "\r\n\r\n");
	pageOrders := [][]int{};
	pageNums := [](map[int]int){};

	for _,v := range strings.Split(data[0], "\r\n") {
		pageOrders = append(pageOrders,sliceAtoi(strings.Split(v, "|")) );
	}
	data[1] = data[1][:len(data[1])-2];
	for i,v := range strings.Split(data[1], "\r\n") {
		t := sliceAtoi(strings.Split(v,","));
		pageNums = append(pageNums, map[int]int{});
		for j,v2 := range t {
			pageNums[i][v2] = j;
			if (j == len(t)/2) {pageNums[i][-1] = v2;}
		}
	}
	// Part_One
	correctMP := 0;
	for _,nMap := range pageNums {
		isCorrect := true;
		for _,order := range pageOrders {
			a,ok1 := nMap[order[0]];
			b,ok2 := nMap[order[1]]
			if ok1 && ok2 && a > b{
				isCorrect = false;
				break;
			}
		}
		if isCorrect {
			correctMP += nMap[-1];
		}
	}
	fmt.Println("Part_One _>",correctMP);
	// Part_Two
	incorrectMP := 0;
	i := 0;
	last := -1;
	for i<len(pageNums){
		nMap := pageNums[i];
		isCorrect := true;
		for _,order := range pageOrders {
			a,ok1 := nMap[order[0]];
			b,ok2 := nMap[order[1]]
			if ok1 && ok2 && a > b{
				isCorrect = false;
				nMap[order[0]] = b;
				nMap[order[1]] = a;
				if nMap[-1] == order[0] {
					nMap[-1] = order[1];
				} else if nMap[-1] == order[1] {
					nMap[-1] = order[0];
				}
				break;
			}
		}
		if !isCorrect {
			last = i;
			continue;
		}
		if last == i {
			incorrectMP += nMap[-1];
		}
		last = i;
		i++;
	}

	fmt.Println("Part_Two _>",incorrectMP)
}
