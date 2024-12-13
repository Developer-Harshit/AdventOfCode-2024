package main

import (
	"fmt"
	"strings"
)
func isSafe(n []int) bool {

	isIncreasing := n[0] < n[1];
	safe := true;
	for j := 0; j < len(n)-1; j++ {
			
		if isIncreasing {
			
			diff := n[j+1] - n[j];
			if !(diff > 0 && diff < 4) {
				safe = false;
				break;
			}
		} else {
				
			diff := (n[j] - n[j+1]);
			if !(diff > 0 && diff < 4) {
				safe = false;
				break;
			}
		}
	}
	return safe;
}
func main(){
	defer timer("program")();
	// Input_Parsing
	data := [][]int{};
	for v := range  readLine() {
		data = append(data,sliceAtoi(strings.Fields(v)));
	}
	// Part_One 
	safeCount := 0;
	for i := 0; i < len(data); i++ {

		if isSafe(data[i]) {
			safeCount++;
		}
	}
	fmt.Println("Part_One _>",safeCount);
	// Part_Two - doing brute force for now
	dampenSafeCount := 0;
	for i := 0; i < len(data); i++ {
		
		// if data is safe without any removal		
		if isSafe(data[i]) {
			dampenSafeCount++;
			continue;
		}
		// else try to remove one element at a time and check the same
		for k := 0; k < len(data[i]); k++ {

			copied := make([]int,len(data[i]));
			copy(copied,data[i]);
			n := append(copied[:k], copied[k+1:]...);
			if isSafe(n) {
				dampenSafeCount++;
				break;
			}
		}
	}
	
	fmt.Println("Part_Two _>",dampenSafeCount);
}
