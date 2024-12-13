package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)


type Disk struct {
	idx int 
	count int
}

func FindChecksum(data []int) int {
	checksum := 0
	for i,v := range data {
		if v == -1 {continue}
		checksum += i * int(v)
	}
	return checksum
}
func FindChecksumBlock(data []Disk) int {
	checksum := 0; idx := 0
	for _,v := range data {
		if v.idx == -1 {idx+=v.count;continue;}
		for j := 0; j < v.count; j++ {
			checksum += v.idx * idx
			idx++
		}
	}
	return checksum
}
func main(){
	defer timer("program")();
	// Input_Parsing
	data := []int{}
	dataBlocks := []Disk{}
	for i,v:= range  strings.TrimRight(readRaw(), "\r\n") {
		val,err := strconv.Atoi(string(v))
		check(err)
		idx := -1
		if i%2==0 {idx = i/2}
		dataBlocks = append(dataBlocks, Disk{idx: idx,count: val})
		for j:=0; j<val; j++ {data = append(data, idx)}
	}
	// Part_One
	left := 0; right := len(data) - 1
	for left < right {
		for ; left < right && data[left] != -1; left++ {}
		for ; right > left && data[right] == -1; right-- {} 
		data[left] = data[right]
		data[right] = -1
	}
	fmt.Println("Part_One _>",FindChecksum(data));
	// Part_Two
	for i := len(dataBlocks) - 1; i > -1; i-- {
		idx1 := dataBlocks[i].idx; c1 := dataBlocks[i].count
		if idx1 == -1 {continue}
		for j := 0; j < i; j++ {
			idx2 := dataBlocks[j].idx; c2 := dataBlocks[j].count
			if idx2 != -1 || c2 < c1 {continue}
			if c2 > c1 {
				dataBlocks[j].count = c2 - c1
				dataBlocks[i].idx = -1
				dataBlocks = slices.Insert(dataBlocks, j, Disk{idx: idx1,count: c1})
			} else {
				dataBlocks[i].idx = idx2
				dataBlocks[j].idx = idx1
			}
			break
		}
	}
	fmt.Println("Part_Two _>",FindChecksumBlock(dataBlocks))
}
