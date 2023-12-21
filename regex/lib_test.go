package regex

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_Match(t *testing.T) {
	testCases := []struct {
		str     string
		pattern string
		match   bool
	}{
		{
			str:     "ab",
			pattern: "ab",
			match:   true,
		},
		{
			str:     "abc",
			pattern: "ab",
			match:   false,
		},
		{
			str:     "a",
			pattern: "a",
			match:   true,
		},
		{
			str:     "abbc",
			pattern: "a(bb)+c",
			match:   true,
		},
		{
			str:     "abbbc",
			pattern: "a(bb)+c",
			match:   false,
		},
		{
			str:     "acc",
			pattern: "a(b|c)c",
			match:   true,
		},
		{
			str:     "abbbbc",
			pattern: "ab+c",
			match:   true,
		},
		{
			str:     "ac",
			pattern: "ab+c",
			match:   false,
		},
		{
			str:     "bcb",
			pattern: "(b|c)+",
			match:   true,
		},
		{
			str:     "abbbbc",
			pattern: "ab*c",
			match:   true,
		},
		{
			str:     "ac",
			pattern: "ab*c",
			match:   true,
		},
		{
			str:     "bcb",
			pattern: "(b|c)*",
			match:   true,
		},
		{
			str:     genStr(20),
			pattern: genPat(20),
			match:   true,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s@%s", tc.str, tc.pattern), func(t *testing.T) {
			if Match(tc.str, tc.pattern) != tc.match {
				if tc.match {
					t.Fatalf("%q should match %q", tc.str, tc.pattern)
				} else {
					t.Fatalf("%q should not match %q", tc.str, tc.pattern)
				}
			}
		})
	}
}

func genStr(n int) string {
	return strings.Repeat("a", n)
}

func genPat(n int) string {
	return strings.Repeat("a?", n) + strings.Repeat("a", n)
}

func BenchmarkMatch(b *testing.B) {
	n := 30
	s := genStr(n)
	p := genPat(n)
	b.Run("BuiltInRegex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			r, err := regexp.MatchString(p, s)
			if err != nil || !r {
				b.Fatal("Should match")
			}
		}
	})

	b.Run("LocalRegex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if !Match(s, p) {
				b.Fatalf("Should match")
			}
		}
	})
}
