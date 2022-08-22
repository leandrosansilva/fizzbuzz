package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

var numbers = []byte("0123456789")

const maxIntLen = len("9223372036854775807")

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func writeLiteral(lastLimit, l int, out io.Writer) (int, error) {
	a := [maxIntLen]int{}

	for i := 0; l != 0; i++ {
		v := l / 10
		r := l % 10
		a[i] = r
		l = v
	}

	limit := min(lastLimit+2, maxIntLen)

	for {
		limit--

		if a[limit] != 0 {
			break
		}
	}

	for i := limit; i >= 0; i-- {
		r := a[i]
		if _, err := out.Write(numbers[r : r+1]); err != nil {
			return 0, err
		}
	}

	return limit, nil
}

var fizzBuzz = []byte("FizzBuzz")

func FizzBuzz(start, end int, sep []byte, out io.Writer) error {
	var (
		limit = maxIntLen
		err   error
		buff  = bytes.NewBuffer(make([]byte, maxIntLen))
	)

	for i := start; i <= end; i++ {
		buff.Reset()

		sliceBegin, sliceEnd := 4, 4

		if i%3 == 0 {
			sliceBegin = 0
		}

		if i%5 == 0 {
			sliceEnd = len(fizzBuzz)
		}

		switch {
		case sliceBegin == sliceEnd:
			if limit, err = writeLiteral(limit, i, buff); err != nil {
				return err
			}
		default:
			if _, err = buff.Write(fizzBuzz[sliceBegin:sliceEnd]); err != nil {
				return err
			}
		}

		if i < end {
			if _, err := buff.Write(sep); err != nil {
				return err
			}
		}

		if _, err := io.Copy(out, buff); err != nil {
			return err
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

	sep := []byte("\n")

	if err := FizzBuzz(1, v, sep, os.Stdout); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
