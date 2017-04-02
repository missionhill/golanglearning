package modules

import (
	"math/big"
	"reflect"
)


func checkConst() {
	conInt := 1
	AssertEqual(reflect.TypeOf(conInt).String(), "int")
	conFloat64 := float64(0.1)
	AssertEqual(reflect.TypeOf(conFloat64).String(), "float64")
	conString := "golang"
	AssertEqual(reflect.TypeOf(conString).String(), "string")
	conBool := false
	AssertEqual(reflect.TypeOf(conBool).String(), "bool")
	conComplex64 := complex64(3+2.5i)
	AssertEqual(reflect.TypeOf(conComplex64).String(), "complex64")
	conByte := byte(1)
	AssertEqual(reflect.TypeOf(conByte).String(), "uint8")
	conRune := rune('a')
	AssertEqual(reflect.TypeOf(conRune).String(), "int32")
}

func checkShortDeclaration(){
	varInt := 1
	AssertEqual(reflect.TypeOf(varInt).String(), "int")
	varFloat64 := float64(0.1)
	AssertEqual(reflect.TypeOf(varFloat64).String(), "float64")
	varString := "golang"
	AssertEqual(reflect.TypeOf(varString).String(), "string")
	varBool := false
	AssertEqual(reflect.TypeOf(varBool).String(), "bool")
	varComplex64 := complex64(3+2.5i)
	AssertEqual(reflect.TypeOf(varComplex64).String(), "complex64")
	varByte := byte(1)
	AssertEqual(reflect.TypeOf(varByte).String(), "uint8")
	varRune := rune('a')
	AssertEqual(reflect.TypeOf(varRune).String(), "int32")

	varIntAddress := &varInt
	// varInt can be re-declared here as long as there is un-declare variable on the left side.
	// varInt := 2 is invalid though.
	varInt, varX := 2, "test"
	AssertEqual(reflect.TypeOf(varX).String(), "string")
	AssertEqual(varIntAddress, &varInt)
	AssertEqual(varInt, 2)
}

func checkGeneralDeclaration(){
	var varInt int
	AssertEqual(reflect.TypeOf(varInt).String(), "int")
	AssertEqual(varInt, 0)
	var varFloat64 float64
	AssertEqual(reflect.TypeOf(varFloat64).String(), "float64")
	AssertTrue(big.NewFloat(0).Cmp(big.NewFloat(varFloat64)) == 0)
	var varString string
	AssertEqual(reflect.TypeOf(varString).String(), "string")
	AssertEqual(varString,"")
	var varBool bool
	AssertEqual(reflect.TypeOf(varBool).String(), "bool")
	AssertEqual(varBool, false)
	var varComplex64 complex64
	AssertEqual(reflect.TypeOf(varComplex64).String(), "complex64")
	AssertEqual(varComplex64, complex64(0+0i))
	var varByte byte
	AssertEqual(reflect.TypeOf(varByte).String(), "uint8")
	AssertEqual(varByte, byte(0))
	var varRune rune
	AssertEqual(reflect.TypeOf(varRune).String(), "int32")
	AssertEqual(varRune, int32(0))

	varInt, varFloat64, varString, varBool, varComplex64, varByte, varRune =
	1, 1.0, "golang", true, 1+1i, byte(1), rune('a')
	AssertEqual(varInt, 1)
	AssertTrue(big.NewFloat(1.0).Cmp(big.NewFloat(varFloat64)) == 0)
	AssertEqual(varString, "golang")
	AssertEqual(varBool, true)
	AssertEqual(varComplex64, complex64(1+1i))
	AssertEqual(varByte, byte(1))
	AssertEqual(varRune, rune('a'))

	var varIntArray [10]int
	AssertEqual(reflect.TypeOf(varIntArray).String(), "[10]int")
	for _, i := range varIntArray {
		AssertEqual(i, 0)
	}

	var varStringArray [10]string
	AssertEqual(reflect.TypeOf(varStringArray).String(), "[10]string")
	for _, i := range varStringArray {
		AssertEqual(i, "")
	}
}

func checkNew(){
	varInt := new(int)
	varInt2 := varInt
	AssertEqual(*varInt, 0)
	*varInt2 = 1
	AssertEqual(*varInt, 1)
}

func checkTypeDeclaration(){
	type Celsius int
	type Fahrenheit int
	var x Celsius
	x = 38
	var y Fahrenheit
	y = 96
	AssertEqual(reflect.TypeOf(x).Name(), "Celsius")
	AssertEqual(reflect.TypeOf(y).Name(), "Fahrenheit")
	AssertTrue(x==38)
	AssertTrue(y==96)
}

func checkScope()  {
	x := "hello"
	AssertEqual(reflect.TypeOf(x).String(), "string")
	for _, x := range x {
		AssertEqual(reflect.TypeOf(x).String(), "int32")
		x := x + 'A' - 'a'
		AssertEqual(reflect.TypeOf(x).String(), "int32")
	}
	AssertEqual(x, "hello")
	if x := "world"; x== "test" {
		AssertEqual(x, "test")
	} else if y:=x; y == "world" { // nested in the first if.
		AssertEqual(y, "world")
	}
}

func ProgramStructureMain()  {
	// const
	checkConst()

	// short variable declaration
	checkShortDeclaration()

	// general variable declaration
	checkGeneralDeclaration()

	// new keyword
	checkNew()

	// type declaration
	checkTypeDeclaration()

	// scope
	checkScope()
}
