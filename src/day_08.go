package main

import (
	"fmt"
)
func displayVisited(d []string,v map[int]bool) {
	w := len(d[0])
	for j,jv := range d {
		for i,iv := range jv {
			_,ok := v[i + j*w]
			if ok && iv == '.' {
				fmt.Printf("%c",'#')
			}			else {
				fmt.Printf("%c",iv)
			}
		}
		fmt.Println()
	}
}
func main(){
	defer timer("program")();
	// Input_Parsing
	data := []string{};
	signalMap := map[rune]([]int){};
	i := 0;
	for v := range readLine() {
		data = append(data, v);
		for j,c := range v {
			if c == '.' {continue;}
			signalMap[c] = append(signalMap[c], i + j * len(v));
		}
		i++;
	}
	// Part_One
	visited := map[int]bool{};
	for _,v := range signalMap {
		for a := 0; a < len(v); a++ {
			var ax int = v[a] / len(data[0]);
			var ay int = v[a] % len(data);
			for b := a+1; b < len(v); b++ {
				var bx int = v[b] / len(data[0]);
				var by int = v[b] % len(data);
				dx := ax - bx; dy := ay - by;
				mx := ax + dx; my := ay + dy;
				nx := bx - dx; ny := by - dy;
				if mx > -1 && my > -1 && mx < len(data[0]) && my < len(data) {
					visited[(mx) + (my) * len(data[0])] = true;
				}
 				if nx > -1 && ny > -1 && nx < len(data[0]) && ny < len(data) {
					visited[(nx) + (ny) * len(data[0])] = true;
				}
			}
		}
	}
	fmt.Println("Part_One _>",len(visited));
	// Part_Two
	visited = map[int]bool{};
	for _,v := range signalMap {
		for a := 0; a < len(v); a++ {
			var ax int = v[a] / len(data[0]);
			var ay int = v[a] % len(data);
			for b := a+1; b < len(v); b++ {
				var bx int = v[b] / len(data[0]);
				var by int = v[b] % len(data);
				dx := ax - bx; dy := ay - by;
				for i:=0;i<len(data[0]);i++ {
					mx := ax + dx*i; my := ay + dy*i;
					nx := bx - dx*i; ny := by - dy*i;
					counter := 0
					if mx > -1 && my > -1 && mx < len(data[0]) && my < len(data) {
						visited[(mx) + (my) * len(data[0])] = true;
					} else {counter++}
	 				if nx > -1 && ny > -1 && nx < len(data[0]) && ny < len(data) {
						visited[(nx) + (ny) * len(data[0])] = true;
					} else {counter++}
					if counter == 2 {break;}
				}
			}
		}
	}
	fmt.Println("Part_Two _>",len(visited));
}

