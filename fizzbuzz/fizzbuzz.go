package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(buildOutputString(i))
	}
}

func buildOutputString(i int) string {
	switch {
	case i%15 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return strconv.FormatInt(int64(i), 10)
	}
}
