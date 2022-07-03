package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		output := strconv.FormatInt(int64(i), 10)
		if i%15 == 0 {
			output = "FizzBuzz"
		} else if i%3 == 0 {
			output = "Fizz"
		} else if i%5 == 0 {
			output = "Buzz"
		}

		fmt.Println(output)
	}
}
