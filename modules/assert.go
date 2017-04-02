package modules

import (
	"os"
	"fmt"
)

func AssertTrue(condition bool){
	if !condition {
		os.Exit(1)
	}
}

func AssertEqual(result interface{}, expected interface{}){
	if result != expected {
		fmt.Printf("Expected %s but got %s\n", expected, result)
	}
}
