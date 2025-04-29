//https://leetcode.com/problems/print-foobar-alternately/
//#concurrency

type FooBar struct {
	n int
	c chan int
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n: n,
		c: make(chan int),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.c <- 0
		<-fb.c
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.c
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.c <- 0
	}
}