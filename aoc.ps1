param ($cmd,$day,$p3)
$filePath = "src/day_$($day).go"
$initialCode = 'package main 
import (
	"fmt"
)
func main(){
	defer timer("program")()
	// Input_Parsing
	data := readRaw()
	// Part_One
	fmt.Println("Part_One _>",data)
	// Part_Two
	fmt.Println("Part_Two _>",data)
}'

if ([int]$day -lt 10){
	$filePath = "src/day_0$($day).go"
}
if ($cmd -eq "new"){
	if ($p3 -eq "f"){
		echo $initialCode > $filePath
	} else {
    	New-Item $filePath -type file && echo $initialCode > $filePath
	}
	vim $filePath
} elseif ($cmd -eq "day") {
	if ($p3 -eq "test"){
		go build -o main.exe $filePath src/common.go && ./main.exe test
	}
	 elseif ($p3 -eq "real") {
	} elseif ($p3 -eq "all"){
		go build -o main.exe $filePath src/common.go && echo "Running Test_> " &&./main.exe test && echo "Running Real_> " && ./main.exe
	} else {
		go build -o main.exe $filePath src/common.go && ./main.exe
	}
} elseif ($cmd -eq "clear") {
	echo input/real.txt > ""
	echo input/test.txt > ""
	Remove-Item main.exe -ErrorAction SilentlyContinue
}
