package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Convert(t *testing.T) {
	testCaseNormal := map[string]struct {
		num    int
		expect string
	}{
		"1を渡すと文字列1を返す": {num: 1, expect: "1"},
	}
	testCaseMulti3 := map[string]struct {
		num    int
		expect string
	}{
		"3を渡すと文字列Fizzを返す": {num: 3, expect: "Fizz"},
	}
	testCaseMulti5 := map[string]struct {
		num    int
		expect string
	}{
		"5を渡すと文字列Buzzを返す": {num: 5, expect: "Buzz"},
	}
	testCaseMulti3And5 := map[string]struct {
		num    int
		expect string
	}{
		"15を渡すと文字列FizzBuzzを返す": {num: 15, expect: "FizzBuzz"},
	}

	t.Run("3の倍数のときは数の変わりにFizzに変換する", func(t *testing.T) {
		for k, v := range testCaseMulti3 {
			t.Run(k, func(t *testing.T) {
				fizzbuzz := NewFizzBuzz()
				assert.Equal(t, v.expect, fizzbuzz.Convert(v.num))
			})
		}
	})
	t.Run("5の倍数のときは数の変わりにFizzに変換する", func(t *testing.T) {
		for k, v := range testCaseMulti5 {
			t.Run(k, func(t *testing.T) {
				fizzbuzz := NewFizzBuzz()
				assert.Equal(t, v.expect, fizzbuzz.Convert(v.num))
			})
		}
	})
	t.Run("3と5両方の倍数のときは数の変わりにFizzBuzzに変換する", func(t *testing.T) {
		for k, v := range testCaseMulti3And5 {
			t.Run(k, func(t *testing.T) {
				fizzbuzz := NewFizzBuzz()
				assert.Equal(t, v.expect, fizzbuzz.Convert(v.num))
			})
		}
	})
	t.Run("その他の数のときはそのまま文字列に変換する", func(t *testing.T) {
		for k, v := range testCaseNormal {
			t.Run(k, func(t *testing.T) {
				fizzbuzz := NewFizzBuzz()
				assert.Equal(t, v.expect, fizzbuzz.Convert(v.num))
			})
		}
	})

}
