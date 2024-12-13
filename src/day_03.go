package main

import (
	"fmt"
	"regexp"
)
func main(){
	defer timer("program")();
	// Input_Parsing
	data := readRaw();
	re := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)");
	re2 := regexp.MustCompile("(?i)[0-9]+");
	re3 := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)|(?i)don't\\(\\)|(?i)do\\(\\)");

	// Part_One
	multCmds := re.FindAllString(data,-1);
	sop := 0
	for _,cmd := range multCmds {
		
		nums := sliceAtoi(re2.FindAllString(cmd,2));
		sop += nums[0]*nums[1];
	}
	fmt.Println("Part_One _>",sop);
	// Part_Two
	multCmds = re3.FindAllString(data,-1);
	sop =  0;
	enabled := true;
	for _,cmd := range multCmds {
	
		if cmd == "do()" {
			enabled = true;
		} else if cmd == "don't()" {
			enabled = false;
		} else if enabled {
			
			nums := sliceAtoi(re2.FindAllString(cmd, 2));
			sop += nums[0] * nums[1];
		}
	}

	fmt.Println("Part_Two _>",sop);
}
