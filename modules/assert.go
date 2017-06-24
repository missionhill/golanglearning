package modules

import (
	"os"
	"fmt"
	"runtime"
)

func AssertTrue(condition bool){
	if !condition {
		_, fileName, fileLine, _ := runtime.Caller(1)
		fmt.Printf("Assertion failed at %d in %s\n", fileLine, fileName)
		os.Exit(1)
	}
}

func AssertEqual(result interface{}, expected interface{}){
	if result != expected {
		fmt.Printf("Expected %s but got %s\n", expected, result)
	}
}
