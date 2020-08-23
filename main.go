package main

import (
	"fmt"

	"github.com/yurakawa/procon-go/FizzBuzz/fizzbuzz"
)

func main() {
	fb := fizzbuzz.NewFizzBuzz()
	for i := 1; i <= 100; i++ {
		fmt.Println(fb.Convert(i))
	}
}
