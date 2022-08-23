package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"

	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
)

var numbers = []byte("0123456789")

const maxIntLen = len("9223372036854775807")

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func writeLiteral(leftLimit, l int, out io.Writer, literalBuffer []byte) (int, error) {
	limit := max(leftLimit-2, 0)

	for i := 0; i < maxIntLen; i++ {
		literalBuffer[i] = 0
	}

	for i := 0; l != 0; i++ {
		v := l / 10
		r := l % 10
		literalBuffer[maxIntLen-i-1] = numbers[r]
		l = v
	}

	for {
		limit++

		if literalBuffer[limit] != 0 {
			break
		}
	}

	slice := literalBuffer[limit:maxIntLen]

	if _, err := out.Write(slice); err != nil {
		return 0, err
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
		limit         = 0
		err           error
		buff          = bytes.NewBuffer(make([]byte, 0, 4096))
		i             = start
		literalBuffer = [maxIntLen]byte{} // NOTE: always escapes to the heap when used on bytes.Buffer.Write()
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
			if limit, err = writeLiteral(limit, i, buff, literalBuffer[:]); err != nil {
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

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	sep := []byte("\n")

	if err := FizzBuzz(1, v, sep, os.Stdout); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
