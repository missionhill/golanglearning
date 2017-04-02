package modules

import (
	"math"
	"math/cmplx"
	"unicode/utf8"
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
func TypesMain(){
	checkInt()
	checkFloat()
	checkComplex()
	checkString()
	checkConstants()
}