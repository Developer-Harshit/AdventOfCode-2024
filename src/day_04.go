package main

import (
	"fmt"
	"regexp"
)
func _printErr(k byte){
	if !(k=='S'||k=='A'||k=='M'||k=='X'){
		fmt.Printf("ERROR_> (%c) , (%v)\n",k,k);
	}
}
func _makeStr(a byte,b byte,c byte,d byte) string {
	_printErr(a);
	_printErr(b);
	_printErr(c);
	_printErr(d);
	
	return fmt.Sprintf("%c%c%c%c",a,b,c,d);
}
var _re ,_re2 *regexp.Regexp;
func _findOcc(d string) int {
	return len(_re.FindAllStringSubmatch(d,-1)) + len(_re2.FindAllStringSubmatch(d,-1));
}
func main(){
	defer timer("program")();
	// Input_Parsing
	data := []string{};
	_re = regexp.MustCompile("(?i)XMAS");
	_re2 = regexp.MustCompile("(?i)SAMX")
	for v := range readLine() {
		data = append(data, v);
	}
	n:= len(data);
	// Part_One
	xmasCount := 0;
	for r:=0;r<n;r++ {
		d1 := "";
		d2 := "";
		for c:=0;c<n;c++ {
			d1 += string(data[r][c]);
			d2 += string(data[c][r]);
		}
		xmasCount+= _findOcc(d1);
		xmasCount+= _findOcc(d2);
	}
	// d1
	for c:=0;c<n;c++ {
		ci := c;
		ri := 0;
		d1 := "";
		d2 := "";
		for ci < n && ri < n {
			d1 += string(data[ri][ci]);
			d2 += string(data[ri][n-ci-1]);
			ci++;
			ri++;
		}
		xmasCount+= _findOcc(d1);
		xmasCount+= _findOcc(d2);
	}
	for r:=1;r<n;r++ {
		ri := r;
		ci := 0;
		d1 := "";
		d2 := "";
		for ci < n && ri < n {
			d1 += string(data[ri][ci]);
			d2 += string(data[ri][n-ci-1]);
			ci++;
			ri++;
		}
		xmasCount+= _findOcc(d1);
		xmasCount+= _findOcc(d2);
	}	
	fmt.Printf("Part_One _> %d\n",xmasCount);
	// Part_Two
	xmasCount = 0;
	for r:=0;r<n-2;r++ {
		for c:=0;c<n-2;c++ {
			v := data[r][c];
			if !(v == 'M' || v == 'S') {continue;}
			a := string(data[r][c])+string(data[r+1][c+1])+string(data[r+2][c+2]);
			b := string(data[r+2][c])+string(data[r+1][c+1])+string(data[r][c+2]);
			if (a == "SAM" || a == "MAS") && (b == "SAM" || b == "MAS") {
				xmasCount++;
			}
		}
	}
	fmt.Println("Part_Two _>",xmasCount);
}

