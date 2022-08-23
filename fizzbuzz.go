package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"

	//"net/http"
	//_ "net/http/pprof"
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

func maybeFlush(i, end int, buff *bytes.Buffer, out io.Writer) error {
	hasEnded := i == end

	timeToFlush := i%512 == 0

	if !(timeToFlush || hasEnded) {
		return nil
	}

	if _, err := io.Copy(out, buff); err != nil {
		return err
	}

	buff.Reset()

	return nil
}

var fizzBuzz = []byte("FizzBuzz")

// FizzBuzz writes on out the FizzBuzz values, separated by sep,
// in the interval [start, end].
// In case any error calling out.Write(), the error is bubbled up
// by FizzBuzz
func FizzBuzz(start, end int, sep []byte, out io.Writer) error {
	var (
		limit = maxIntLen
		err   error
		buff  = bytes.NewBuffer(make([]byte, 0, 4096))
		i     = start
	)

	for ; i <= end; i++ {
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

		if err := maybeFlush(i, end, buff, out); err != nil {
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

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	sep := []byte("\n")

	if err := FizzBuzz(1, v, sep, os.Stdout); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
