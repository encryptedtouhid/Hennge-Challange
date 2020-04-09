package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var Totn int
var tot []int

func main() {
	var outerloop int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
		
	outerloop, err := strconv.Atoi(scanner.Text())
    if err != nil{
    	fmt.Println("outerloop error")
    }
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
    Totn = outerloop
	fmt.Println(Totn)
}