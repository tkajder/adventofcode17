/*
 * Copyright 2017 Thomas Kajder
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"container/ring"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func convertToSingleRuneIntRing(s string) (*ring.Ring, error) {
	intRing := ring.New(len(s))
	for _, c := range s {
		if unicode.IsDigit(c) {
			intRing.Value = int(c - '0')
			intRing = intRing.Next()
		} else {
			return nil, errors.New("Non digit character [" + string(c) + "] encountered")
		}
	}

	return intRing, nil
}

func calculateCaptcha(r *ring.Ring) int {
	sum := 0
	for i := 0; i < r.Len(); i++ {
		if matchesNext(r) {
			sum += r.Value.(int)
		}
		r = r.Next()
	}

	return sum
}

func matchesNext(r *ring.Ring) bool {
	return r.Value == r.Next().Value
}

func main() {
	fileNamePtr := flag.String("file", "", "Required: file containing inverse captcha")
	flag.Parse()

	// Check if file provided
	if *fileNamePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(b), " \t\n\r")

	inverseCaptcha, err := convertToSingleRuneIntRing(s)
	if err != nil {
		log.Fatal(err)
	}

	sum := calculateCaptcha(inverseCaptcha)
	fmt.Println(sum)
}
