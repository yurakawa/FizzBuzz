package fizzbuzz

import "strconv"

type FizzBuzz struct{}

func NewFizzBuzz() FizzBuzz {
	return FizzBuzz{}
}

func (f FizzBuzz) Convert(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	}
	if num%3 == 0 {
		return "Fizz"
	}
	if num%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(num)
}
