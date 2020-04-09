package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var totaltestcases int
var tot []int

func main() {

	totaltestcases := testcaseinput()

	fmt.Println("Total Test Cases : " + strconv.Itoa(totaltestcases))
}

func testcaseinput() int {

	var outerloop int
	fmt.Print("Total Number Of Test Cases : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	outerloop, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Not a Valid Number. Please try again")
		testcaseinput()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return outerloop
}
