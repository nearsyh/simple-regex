package regex

import (
	"strings"
	"testing"
)

func Test_Parse(t *testing.T) {
	testCases := []struct{
		pattern string
		expected string
	}{
		{
			pattern: "a",
			expected: "a",
		},
		{
			pattern: "ab",
			expected: "ab.",
		},
		{
			pattern: "abc",
			expected: "ab.c.",
		},
		{
			pattern: "b|c",
			expected: "bc|",
		},
		{
			pattern: "(b|c)",
			expected: "bc|",
		},
		{
			pattern: "a(b|c)+(d|ef)+",
			expected: "abc|+.def.|+.",
		},
		{
			pattern: "(b|c)+",
			expected: "bc|+",
		},
		{
			pattern: "(b|cd)*",
			expected: "bcd.|*",
		},
		{
			pattern: "a?a?a?aaa",
			expected: "a?a?.a?.a.a.a.",
		},
		{
			pattern: genPat(20),
			expected: "a?" + strings.Repeat("a?.", 19) + strings.Repeat("a.", 20),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.pattern, func(t *testing.T) {
			actual := toPost(tc.pattern)
			if actual != tc.expected {
				t.Fatalf("Got %q, want %q", actual, tc.expected)
			}
		})
	}
}