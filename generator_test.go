package angen

import (
	"strings"
	"testing"
)

func TestGenerateNumber(t *testing.T) {
	gen := NewGenerator(3, 4, 0, 10000000)
	ch, closer := gen.Gen()
	var result []string
	defer close(closer)

	for v := range ch {
		result = append(result, *v)
		if len(result) >= 5 {
			break
		}
	}

	t.Logf("size: %d, result:%s", len(result), strings.Join(result, ","))
}

func TestGenerateNumber2(t *testing.T) {
	gen := NewGenerator(1, 1, 0, 1000)
	ch, _ := gen.Gen()
	var result []string

L:
	for {
		select {
		case v := <-ch:
			if v != nil {
				result = append(result, *v)
			} else {
				break L
			}
		}
	}

	t.Logf("size: %d, result:%s", len(result), strings.Join(result, ","))
}

func BenchmarkGenerateNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gen := NewGenerator(1, 3, 0, 10)
		ch, _ := gen.Gen()
		var result []string

	L:
		for {
			select {
			case v := <-ch:
				if v != nil {
					result = append(result, *v)
				} else {
					break L
				}
			}
		}
	}
}
