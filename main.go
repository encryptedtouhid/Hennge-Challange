package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var totaltcs int
var tot []int

func main() {
	totaltcs := testcaseinput()
	fmt.Println("Total Test Cases : " + strconv.Itoa(totaltcs))
}

func testcaseinput() int {
	var totaltestcases int

	fmt.Print("Total Number Of Test Cases : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	totaltestcases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Not a Valid Number. Please try again")
		return testcaseinput()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	return totaltestcases
}
