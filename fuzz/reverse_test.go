package main

import (
	"testing"
	"unicode/utf8"
)


func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, origin string){
		rev, err1 := Reverse(origin)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		t.Logf("number of runes: original=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(origin), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if origin != doubleRev {
			t.Errorf("Before: %q, after: %q", origin, doubleRev)
		}
		if utf8.ValidString(origin) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}