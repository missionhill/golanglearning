package modules

import (
	"math"
	"math/cmplx"
	"unicode/utf8"
	"reflect"
	"unsafe"
)

func checkInt(){
	var varInt32 int32
	varInt32 = 1<<31-1
	AssertTrue(varInt32 > 0)
	AssertTrue(varInt32 + 1 == -(1<<31))
	varInt32 = -(1<<31)
	AssertTrue(varInt32 < 0)
	varInt32 = varInt32 -1
	AssertTrue(varInt32 == 1<<31 -1)
	var varUint32 uint32
	varUint32 = 0
	AssertTrue(varUint32 -1 == 1<<32 - 1)
	varUint32 = 1<<32 -1
	AssertTrue(varUint32+1 == 0)

	varUint32 = 1 << 31
	AssertTrue(varUint32>>1 == 1<<30)
	AssertTrue(varUint32 << 1 == 0)
	varUint32 |= varUint32 -1
	varInt32 = int32(varUint32)
	AssertTrue(varInt32 == -1)
	AssertTrue(varInt32 == varInt32 >>1)
	AssertTrue(varInt32 << 1 == -2)
}

func checkFloat(){
	var varFloat32 float32
	varFloat32 = 1<<30
	AssertTrue(varFloat32+1 == varFloat32)
	var varFloat64 float64
	AssertTrue(varFloat64 == -varFloat64)
	AssertTrue(math.IsInf(1/varFloat64, 1))
	AssertTrue(math.IsNaN(varFloat64/varFloat64))
}

func checkComplex(){
	AssertEqual(cmplx.Sqrt(-1), 0+1i)
	AssertEqual((1+2i)*(2+3i), -4+7i )
}

func checkString(){
	varString := "hello world"
	AssertEqual(varString[:5], "hello")
	AssertEqual(varString[len(varString)-5:], "world")
	varString = "Hello, 世界"
	// Each chinese character needs three chars
	AssertEqual(len(varString), 13)
	varRunes := []rune(varString)
	AssertEqual(len(varRunes), 9)
	AssertEqual(utf8.RuneCountInString(varString), 9)
}

func checkConstants(){
	type Weekday int

	const (
		Sunday Weekday = iota // auto increment
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	AssertEqual(Friday, Weekday(5))

	const (
		One = 1 << iota
		Two
		Four
		Eight
	)
	AssertEqual(Eight, 8)
	const (
		_ = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776
	 	PiB // 1125899906842624
	   	EiB // 1152921504606846976
		ZiB // 1180591620717411303424
	    YiB // 1208925819614629174706176)
	)
	//YiB and ZiB is untyped because they exceeds the limit of integer type.
	AssertEqual(YiB/ZiB, 1024)
	// LARGE_INT is untyped const
	const LARGE_INT=1<<216
	AssertTrue(LARGE_INT > math.MaxInt64)
}

func checkArray(){
	varArray := [...]int{1,2,3}
	AssertEqual(reflect.TypeOf(varArray).String(), "[3]int")

	// array initialization using index
	const (
		_  = iota
		USD
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	AssertEqual(len(symbol), 5)
	AssertEqual(symbol[0], "")

	// array comparison
	x := [...]int{1,2,3}
	y := [...]int{1,2,3}
	AssertEqual(x, y)
	y[2] = 4
	AssertTrue(x != y)
}

func checkSlice() {
	varArray := [10]int{}
	for index, _ := range (varArray) {
		varArray[index] = index
	}
	AssertEqual(len(varArray[1:3]), 2)
	AssertEqual(cap(varArray[1:3]), 9)
	x := []int{1, 2, 3}    // slice
	y := [...]int{1, 2, 3} // array
	reverseArray1(x)
	AssertTrue(x[0] == 3 && x[1] == 2 && x[2] == 1)
	reverseArray2(y)
	AssertTrue(y == [3]int{1, 2, 3})

	var s []int
	AssertTrue(s == nil)
	s = []int(nil)
	AssertTrue(s == nil)
	s = []int{}
	AssertTrue(s != nil)
	var varSlice []int
	varSlice = append(varSlice, 0)
	address1 := &varSlice[0]
	for i := 1 ; i < 5 ; i++ {
		varSlice = append(varSlice,i)
	}
	address2 := &varSlice[0]
	AssertEqual(cap(varSlice), 8)
	AssertTrue(address1!=address2)
	varSlice = append(varSlice, 5)
	AssertEqual(address2, &varSlice[0])

	s = nil
	s = append(s, 1)
	AssertEqual(len(s), 1)

}

func checkMap(){
	var varMap map[string]int
	AssertTrue(varMap == nil)
	value, ok := varMap["first"]
	AssertEqual(value, 0)
	AssertEqual(ok, false)
	varMap = make(map[string]int)
	varMap["first"] = 1
	value, ok = varMap["first"]
	AssertEqual(value, 1)
	AssertEqual(ok, true)
}

func checkStruct(){
	type Empty struct {
	}
	var varEmptyStruct Empty // empty struct occupies no memory for space saving
	AssertEqual(int(unsafe.Sizeof(varEmptyStruct)), 0)

	type Point struct {
		X int
		Y int
	}
	p1 := new(Point)
	*p1 = Point{1,2}
	p2 := &Point{1,2}
	AssertEqual(reflect.TypeOf(p1), reflect.TypeOf(p2))
	AssertTrue(*p1==*p2)
	points := make(map[Point]int)
	points[Point{1,2}] = 1
	value, ok := points[Point{1,2}]
	AssertEqual(value, 1)
	AssertTrue(ok)

	type Circle struct {
		Center Point
		Radius int
	}

	type Wheel struct {
		Circle Circle
		Spokes int
	}

	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Center:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	AssertTrue(w1==w2)
}

func reverseArray1(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArray2(s [3]int){
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TypesMain(){
	checkInt()
	checkFloat()
	checkComplex()
	checkString()
	checkConstants()
	checkArray()
	checkSlice()
	checkMap()
	checkStruct()
}