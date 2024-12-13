package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
func genBinaryString(k int,v int) string {
	str := fmt.Sprintf("%b",v);
	return strings.Repeat("0",k-len(str)) + str;
}
func genTernaryStringList(k int) []string {
	n := int(math.Pow(3,float64(k)));
	res := []string{};
	for j := 0; j < n; j++ {
		res = append(res,"");
	}
	for i := 0; i < k; i++ {
		z := 0;
		t := 0;
		for j := 0; j < n; j++ {
			res[j] += fmt.Sprintf("%d",z);
			t++;
			if t == int(math.Pow(3,float64(i))) {
				t = 0;
				z = (z+1) % 3;
			}
		}
	}
	return res;
}
type Equation struct {
	sum int 
	nums []int
} 
func main(){
	defer timer("program")();
	// Input_Parsing
	data := []Equation{};
	for v := range readLine() {
		tmp := strings.Split(v, ": ");
		s,err := strconv.Atoi(tmp[0]);
		check(err);
		n := sliceAtoi(strings.Split(tmp[1], " "));
		data = append(data, Equation{s,n})
	}
	result := 0;
	for _,eq := range data {
		sum := eq.sum;
		nums := eq.nums;
		n := int(math.Pow(float64(2),float64(len(nums)-1)));
		for v := 0; v < n; v++ {
			bin := genBinaryString(len(nums)-1,v);
			computedSum := nums[0];
			for i,b := range bin {
				
				if b == '0' {
					computedSum += nums[i+1];
				} else if b == '1' {
					computedSum *= nums[i+1];
				}
			}
			if computedSum == sum {
				result += sum;
				break;
			}
		}
	}
	// Part_One
	fmt.Println("Part_One _>",result);
	// Part_Two
	result = 0;
	for _,eq := range data {
		sum := eq.sum;
		nums := eq.nums;
		tenList := genTernaryStringList(len(nums)-1);
		for _,ten := range tenList {
			computedSum := nums[0];
			for i,b := range ten {
				if b == '0' {
					computedSum += nums[i+1];
				} else if b == '1' {
					computedSum *= nums[i+1];
				} else {
					c,_ := strconv.Atoi(fmt.Sprintf("%d%d",computedSum,nums[i+1]));
					computedSum = c;
				}
			}
			if computedSum == sum {
				result += sum;
				break;
			}
		}
	}
	fmt.Println("Part_Two _>",result);
}
