package main

import (
	"io"
)

var numbers = []byte("0123456789")

func writeLiteral(l int, out io.Writer) {
	a := [maxIntLen]int{}

	for i := 0; l != 0; i++ {
		v := l / 10
		r := l % 10

		a[i] = r

		l = v
	}

	limit := maxIntLen

	for {
		limit--

		if a[limit] != 0 {
			break
		}
	}

	for i := limit; i >= 0; i-- {
		r := a[i]
		out.Write(numbers[r : r+1])
	}
}

const maxIntLen = len("9223372036854775807")

var fizzBuzz = []byte("FizzBuzz")

func FizzBuzz(start, end int, out io.Writer) {
	for i := start; i <= end; i++ {
		sliceBegin, sliceEnd := 4, 4

		if i%3 == 0 {
			sliceBegin = 0
		}

		if i%5 == 0 {
			sliceEnd = len(fizzBuzz)
		}

		switch {
		case sliceBegin == sliceEnd:
			writeLiteral(i, out)
		default:
			out.Write(fizzBuzz[sliceBegin:sliceEnd])
		}

		if i < end {
			out.Write([]byte(", "))
		}
	}
}
