package main

import (
	"fmt"
	"testing"
)

/*
goos: darwin
goarch: arm64
pkg: go-mini-quiz
BenchmarkFindWords/4_chars-8         	   77313	     13636 ns/op	    4536 B/op	      96 allocs/op
BenchmarkFindWords/8_chars-8         	   66728	     18092 ns/op	    4608 B/op	     111 allocs/op
BenchmarkFindWords/12_chars-8        	   32774	     36995 ns/op	   28299 B/op	     278 allocs/op
BenchmarkFindWords/50_chars-8        	    6459	    184922 ns/op	  185701 B/op	     990 allocs/op
BenchmarkFindWords/100_chars-8       	    4213	    284033 ns/op	  192122 B/op	    1128 allocs/op
BenchmarkFindWords/200_chars-8       	    2145	    563481 ns/op	  411842 B/op	    1452 allocs/op
PASS
ok  	go-mini-quiz	8.342s
*/
func BenchmarkFindWords(b *testing.B) {
	testLengths := []int{4, 8, 12, 50, 100, 200}
	for _, length := range testLengths {
		str := generateRandomString(length)
		b.Run(fmt.Sprintf("%d chars", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findWords(str)
			}
		})
	}
}
