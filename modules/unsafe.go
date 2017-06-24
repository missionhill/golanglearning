package modules

import (
	"unsafe"
	"strings"
	"reflect"
)

func checkUnsafePtr(){
	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	AssertTrue(x.b == 42)
}

func checkDeepEqual(){
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"};
	AssertTrue(reflect.DeepEqual(got, want))
	var a, b []string = nil, []string{}
	AssertTrue(!reflect.DeepEqual(a,b))
	var c, d map[string]int = nil, make(map[string]int)
	AssertTrue(!reflect.DeepEqual(c,d))
}

func UnsafeMain(){
	checkUnsafePtr()
	checkDeepEqual()
}
