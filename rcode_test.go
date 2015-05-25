package rcode_test

import (
	"fmt"
	"github.com/wridgers/rcode"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGenerate(t *testing.T) {
	var cases = []int{0, 5, 10}

	for _, c := range cases {
		code, err := rcode.Generate(c)

		if err != nil {
			t.Errorf("Generate(%d) returned non nil error value")
		}

		if utf8.RuneCountInString(code) != c {
			t.Errorf("RuneCount of Generate(%d) \"%s\" is not %d", c, code, c)
		}
	}
}

func ExampleGenerate() {
	fmt.Println(rcode.Generate(10))
}

func TestGenerateFromSource(t *testing.T) {
	var cases = []struct {
		source string
		length int
	}{
		{"a", 5},
		{"b", 0},
		{"世", 10},
		{"0", 5},
	}

	for _, c := range cases {
		actualCode, err := rcode.GenerateFromSource(c.source, c.length)

		if err != nil {
			t.Errorf("GenerateFromSource(%s, %d) returned a non nil error value", c.source, c.length)
		}

		expectedCode := strings.Repeat(c.source, c.length)

		if actualCode != expectedCode {
			t.Errorf("GenerateFromSource(%s, %d) \"%s\" is not \"%s\"", c.source, c.length, actualCode, expectedCode)
		}
	}
}

func ExampleGenerateFromSource() {
	fmt.Println(rcode.GenerateFromSource("abcdef", 10))
}

func benchmarkGenerate(length int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		rcode.Generate(length)
	}
}

func BenchmarkGenerate5(b *testing.B)   { benchmarkGenerate(5, b) }
func BenchmarkGenerate10(b *testing.B)  { benchmarkGenerate(10, b) }
func BenchmarkGenerate50(b *testing.B)  { benchmarkGenerate(50, b) }
func BenchmarkGenerate100(b *testing.B) { benchmarkGenerate(100, b) }
