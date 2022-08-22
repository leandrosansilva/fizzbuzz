package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

var numbers = []byte("0123456789")

func writeLiteral(l int, out io.Writer) error {
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
		if _, err := out.Write(numbers[r : r+1]); err != nil {
			return err
		}
	}

	return nil
}

const maxIntLen = len("9223372036854775807")

var fizzBuzz = []byte("FizzBuzz")

func FizzBuzz(start, end int, sep string, out io.Writer) error {
	sepBytes := []byte(sep)

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
			if err := writeLiteral(i, out); err != nil {
				return err
			}
		default:
			if _, err := out.Write(fizzBuzz[sliceBegin:sliceEnd]); err != nil {
				return err
			}
		}

		if i < end {
			if _, err := out.Write(sepBytes); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if v == 0 {
		v = math.MaxInt
	}

	if err := FizzBuzz(1, v, "\n", os.Stdout); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
