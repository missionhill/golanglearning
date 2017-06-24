package modules

import (
	"testing"
	"fmt"
)

func TestCheckDeferredResult(t *testing.T){
	if checkDeferredResult() != 3{
		t.Errorf("Wrong checkDeferredResult")
	}
}

func BenchmarkCheckDeferredResultBenchmark(b *testing.B){
	for i := 0; i < b.N; i++ {
		checkDeferredResult()
	}
}

func Foo() int{
	return 1
}

func ExampleFoo() {
	fmt.Println(Foo())
	//output:
	//1
}