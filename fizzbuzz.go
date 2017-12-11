package fizzbuzz

import "fmt"

func FizzBuzz(times int) string {
	results := ""
	for fizzBuzzCount := 1; fizzBuzzCount <= times; fizzBuzzCount++ {
		if fizzBuzzCount%3 == 0 {
			results += "Fizz\n"
		}
		if fizzBuzzCount%5 == 0 {
			results += "Buzz\n"
		}
		results += fmt.Sprintf("%d\n", fizzBuzzCount)
	}
	return results
}
