package modules

import (
	"reflect"
)

func func1(x,y string) string{
	return "this is func1"
}

func func2(x,y string) string  {
	return "this is func2"
}

func func3(x,y string) (result, error string){
	return x+y, ""
}

func func4(vals ... int) int{
	sum := 0
	for _, val := range(vals){
		sum += val
	}
	return sum
}

func func5(vals []int) int { return func4(vals...)}

func squares() func() int{
	var x int // x's life is not determined by its scope
	return func() int{
		x++
		return x*x
	}
}

func checkSignature(){
	AssertEqual(reflect.TypeOf(func1), reflect.TypeOf(func2))
	result, error := func3("hello", "world")
	AssertEqual(result, "helloworld")
	AssertEqual(error, "")
}

func checkValues(){
	var varFunc func(string, string) string
	AssertTrue(varFunc == nil)
	varFunc = func1
	varResult := varFunc("","")
	AssertEqual(varResult, func1("",""))
	varFunc= func2
	varResult = varFunc("","")
	AssertEqual(varResult, func2("",""))
}

func checkClosure(){
	f := squares()
	AssertEqual(f(), 1)
	AssertEqual(f(), 4)
}

func checkVariadic(){
	AssertTrue(reflect.TypeOf(func4).String() != reflect.TypeOf(func5).String())
	vals := []int{1,2,3,4}
	AssertEqual(func4(vals...), func5(vals) )
	AssertEqual(func5(vals), func4(1,2,3,4))
}

func checkDeferred() int{
	x := 0
	defer func(value int){ // value is evaluated when the statement is executed.
		x++ // x is evaluated when the function is called
		AssertEqual(value, 0)
		AssertEqual(x, 4)
	}(x);
	x++
	defer func(value int){ // value is evaluated when the statement is executed.
		x++ // x is evaluated when the function is called
		AssertEqual(value, 1)
		AssertEqual(x, 3)
	}(x)
	x++
	return x // deferred functions are called after the return
}

func checkDeferredResult() (result int){
	defer func(){
		result = 3
	}()
	return 1
}

func checkPanicRecovery(){
	defer func() {
		if p := recover(); p != nil {
			AssertEqual(p, "panic")
		}
	}()
	panic("panic")
}

func FunctionMain(){
	checkSignature()
	checkValues()
	checkClosure()
	checkVariadic()
	result :=checkDeferred()
	AssertEqual(result, 2)
	result = checkDeferredResult()
	AssertEqual(result, 3)
	checkPanicRecovery()
}
