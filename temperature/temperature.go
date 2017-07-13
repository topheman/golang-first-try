package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	temps := scanner.Text() // the n temperatures expressed as integers ranging from -273 to 5526
	numbers := strings.SplitN(temps, " ", n)
	// fmt.Println(temps)
	// fmt.Println(numbers)

	var result float64 = 5526
	if n == 0 {
		result = 0
	}
	if n == 1 {
		result, _ = strconv.ParseFloat(numbers[0], 64)
	}
	if n > 1 {
		var i = 0
		for i = 0; i < n; i++ {
			parsed, _ := strconv.ParseFloat(numbers[i], 64)
			if math.Abs(parsed) < math.Abs(result) {
				result = parsed
			}
			if math.Abs(parsed) == math.Abs(result) && parsed > 0 {
				result = parsed
			}
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(result) // Write answer to stdout
}

// https://www.codingame.com/ide/puzzle/temperatures
// run:
// cat temperature/fixtures | go run temperature/temperature.go
