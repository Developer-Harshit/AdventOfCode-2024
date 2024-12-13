package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
	"time"
)
func check(e error) {
	if e != nil {
		panic(e);
	}
}
func timer(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}
func getFilePath() string {
	args := os.Args;
	fmt.Println(args);
	if len(args) < 2 || args[1] == "real" {
		return "./input/real.txt";
	}
	return "./input/test.txt";
}
func readRaw() string {
	data , err := os.ReadFile(getFilePath());
	check(err);
	return string(data[:]);
}
func readLine() iter.Seq[string] {
	
	return func(yield func(string) bool){

		inputFilePath := getFilePath();
		f,err := os.Open(inputFilePath);
		defer f.Close();
		check(err);
		scanner := bufio.NewScanner(f);
		for scanner.Scan() {
		
			yield(scanner.Text());
		}
		err = scanner.Err();
		check(err);
		return;
	}
}

func sliceAtoi(arr []string) []int {

	iArr := make([]int,0,len(arr));
	for _,a := range arr {

		i,err := strconv.Atoi(a);
		check(err);
		iArr = append(iArr,i);
	}
	return iArr;
}

