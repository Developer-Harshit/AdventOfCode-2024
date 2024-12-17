package main

import (
	"fmt"
	"regexp"
)
type MachineData []float64
func main(){
	defer timer("program")()
	// Input_Parsing
	data := []MachineData{}
	curr := MachineData{}
	i := 0
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString("abc123def987asdf", -1))
	for v := range readLine() {
		if v == "" {continue}
		nums := sliceAtoi(re.FindAllString(v,-1))
		if len(nums) != 2 {fmt.Errorf("Invalid input")}
		curr = append(curr, float64(nums[0]))		
		curr = append(curr, float64(nums[1]))
		if i % 3 == 2 {
			data = append(data, curr)
			curr = MachineData{}
		}
		i++
	}
	// Part_One
	tokenCostSum := 0
	for _,v := range data {
		x := (v[5] * v[2] - v[4] * v[3]) / (v[2] * v[1] - v[0] * v[3])
		y := (v[4] - x * v[0]) / v[2]
		if !(float64(int(x)) == x && float64(int(y)) == y) { continue }
		if x < 0 || y < 0 || x > 100 || y > 100 { continue }
		tokenCostSum += int(x) * 3 + int(y) * 1
	}
	fmt.Printf("Part_One _> %#v\n",tokenCostSum)
	// Part_Two
	k := float64 (10000000000000)
	tokenCostSum = 0
	for _,v := range data {
		x := ((v[5]+k) * v[2] - (v[4]+k) * v[3]) / (v[2] * v[1] - v[0] * v[3])
		y := ((v[4]+k) - x * v[0]) / v[2]
		if !(float64(int(x)) == x && float64(int(y)) == y) { continue }
		if x < 0 || y < 0 { continue }
		tokenCostSum += int(x) * 3 + int(y) * 1
	}
	fmt.Println("Part_Two _>",tokenCostSum)
}
