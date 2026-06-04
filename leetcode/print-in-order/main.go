// https://leetcode.com/problems/print-in-order/

package main

type Foo struct {
	done1 chan struct{}
	done2 chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		done1: make(chan struct{}),
		done2: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	close(f.done1)
}

func (f *Foo) Second(printSecond func()) {
	<-f.done1
	printSecond()
	close(f.done2)
}

func (f *Foo) Third(printThird func()) {
	<-f.done2
	printThird()
}
