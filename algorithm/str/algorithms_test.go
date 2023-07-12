package str

import (
	"strings"
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestZfunction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		line      string
		zfunction []int
	}{
		{"aaaAAA", []int{0, 2, 1, 0, 0, 0}},
		{"abacaba", []int{0, 0, 1, 0, 3, 0, 1}},
		{"babaabbb", []int{0, 0, 2, 0, 0, 1, 1, 1}},
	}
	for i := range tests {
		ans := ZFunction(tests[i].line)
		if !slices.Equal(ans, tests[i].zfunction) {
			t.Errorf("test %d wrong answer for line %s", i+1, tests[i].line)
			t.Errorf("%v %v", ans, tests[i].zfunction)
		}
	}
}

func TestPrefixFunction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		line   string
		prefix []int
	}{
		{"abcdabcabcdabcdab", []int{0, 0, 0, 0, 1, 2, 3, 1, 2, 3, 4, 5, 6, 7, 4, 5, 6}},
		{"abcdabscabcdabia", []int{0, 0, 0, 0, 1, 2, 0, 0, 1, 2, 3, 4, 5, 6, 0, 1}},
	}
	for i := range tests {
		ans := PrefixFunction(tests[i].line)
		if !slices.Equal(ans, tests[i].prefix) {
			t.Errorf("test %d wrong answer for line %s", i+1, tests[i].line)
			t.Errorf("%v %v", ans, tests[i].prefix)
		}
	}
}

func FuzzZFunction(f *testing.F) {
	f.Add("zfunction")
	f.Fuzz(func(t *testing.T, s string) {
		zfunction := ZFunction(s)
		for i, e := range zfunction {
			if i == 0 {
				continue
			}
			if !strings.HasPrefix(s, s[i:i+e]) {
				t.Errorf("wrong answer %s %v", s, zfunction)
			}
			if len(s) > i+e && strings.HasPrefix(s, s[i:i+e+1]) {
				t.Errorf("wrong answer %s %v", s, zfunction)
			}
		}
	})
}

func FuzzPrefixFunction(f *testing.F) {
	f.Add("prefixfunction")
	f.Fuzz(func(t *testing.T, s string) {
		prefixfunction := PrefixFunction(s)
		for i, e := range prefixfunction {
			if i == 0 {
				continue
			}
			if !strings.HasPrefix(s, s[i+1-e:i+1]) {
				t.Errorf("wrong answer %s %v", s, prefixfunction)
			}
			if i != e && strings.HasPrefix(s, s[i-e:i+1]) {
				t.Errorf("wrong answer %s %v", s, prefixfunction)
			}
		}
	})
}
