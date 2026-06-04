package main

import (
	"bytes"
	"sync"
	"testing"
)

func TestFooOrder(t *testing.T) {
	tests := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	for _, order := range tests {
		var buf bytes.Buffer
		foo := NewFoo()
		var wg sync.WaitGroup
		wg.Add(3)
		funcs := []func(){
			func() { buf.WriteString("first"); wg.Done() },
			func() { buf.WriteString("second"); wg.Done() },
			func() { buf.WriteString("third"); wg.Done() },
		}
		go func(i int) { foo.First(funcs[0]) }(order[0])
		go func(i int) { foo.Second(funcs[1]) }(order[1])
		go func(i int) { foo.Third(funcs[2]) }(order[2])
		wg.Wait()
		if buf.String() != "firstsecondthird" {
			t.Errorf("order %v: got %q, want %q", order, buf.String(), "firstsecondthird")
		}
	}
}
