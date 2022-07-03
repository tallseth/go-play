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
	switch 0 {
	case i % 15:
		return "FizzBuzz"
	case i % 3:
		return "Fizz"
	case i % 5:
		return "Buzz"
	default:
		return strconv.FormatInt(int64(i), 10)
	}
}
