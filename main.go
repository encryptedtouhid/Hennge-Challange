package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var totaltcs int
var numberofint int
var inputs []int
var index int

func main() {
	// totaltcs := testcaseinput()
	// numberofint := inputount()
	// fmt.Println("Total Test Cases : " + strconv.Itoa(totaltcs))
	// fmt.Println("Total Number of Inputs : " + strconv.Itoa(numberofint))

	// fmt.Println("Inputs : " + getinput())
	testArray := strings.Fields(getinput())
	stringtoarray(testArray, index)
	fmt.Println(inputs)
}

func stringtoarray(arr []string, loc int) {
	if loc < len(arr) {
		val, err := strconv.Atoi(arr[loc])
		if err == nil {
			// fmt.Print(val)
			index = index + 1
			inputs = append(inputs, val)
			stringtoarray(arr, index)
		}
	}
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

func inputount() int {

	var inputcount int

	fmt.Print("Total Number Of Input : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	inputcount, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Not a Valid Number. Please try again")
		return inputount()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
	return inputcount
}

func getinput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}
