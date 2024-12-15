package main

import (
	"fmt"
	"strconv"
	"strings"
)
func countCacheLen(data BlinkCache) int {
	counter := 0 
	for _,c := range data {counter += c}
	return counter
}
type BlinkCache map[int]int
func blink(data BlinkCache, maxCount int) BlinkCache {
	for i := 0; i < maxCount; i++ {
		tmp := BlinkCache{}
		for k,c := range data {
			if c == 0 {delete(data,k);continue}
			var res []int;
			sval := fmt.Sprintf("%d", k); 
			n := len(sval)
			if	k == 0 {							// replacing number with 1 if its 0 
				res = []int{1}
			} else if n % 2 == 0 {					// spliting numbers if its digits has even count 
				v1, err := strconv.Atoi(sval[:n/2])
				check(err)
				v2, err := strconv.Atoi(sval[n/2:])
				check(err)
				res = []int{v1,v2}
			} else {								// else just multiply the number by 2024
				res = []int{k*2024}
			}
			delete(data,k)
			for _,k2 := range res {
				_,ok := tmp[k2]
				if !ok {tmp[k2] = 0}
				tmp[k2] += c
			}
		}
		for k,c := range tmp {
			_,ok := data[k]
			if !ok {data[k] = 0}
			data[k] += c
		}
	}
	return data
}
func main(){
	defer timer("program")()
	// Input_Parsing
	data := BlinkCache{}
	for _,k := range sliceAtoi(strings.Split(strings.TrimRight(readRaw(), "\r\n"), " ")) {
		data[k] = 1									// assuming the input data has unique elements
	}
	// Part_One
	data = blink(data,25)
	fmt.Println("Part_One _>", countCacheLen(data))
	data = blink(data,50)
	// Part_Two
	fmt.Println("Part_Two _>",countCacheLen(data))
}
