package modules

import (
	"fmt"
	"reflect"
	"bytes"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func checkIOWriteInterface(){
	var c ByteCounter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	AssertEqual(int(c), 12)
}

type Interface1 interface {
	func1 (int) string
}

type Interface2 interface {
	func2 (string) int
}

type Interface3 interface{
	Interface1
	Interface2
}

type Type1 struct{}
func (t Type1) func1(int) string{
	return "type1"
}
type Type2 struct{}
func (t Type2) func2(string) int{
	return 1
}
type Type3 struct{}
func (t Type3) func1(int) string{
	return "type1"
}
func (t Type3) func2(string) int{
	return 1
}

func checkInterfaceImplementations(){
	t1 := Type1{}
	i1 := reflect.TypeOf((*Interface1)(nil)).Elem()
	AssertTrue(reflect.TypeOf(t1).Implements(i1))
	t2 := Type2{}
	i2 := reflect.TypeOf((*Interface2)(nil)).Elem()
	AssertTrue(reflect.TypeOf(t2).Implements(i2))
	t3 := Type3{}
	i3 := reflect.TypeOf((*Interface3)(nil)).Elem()
	AssertTrue(reflect.TypeOf(t3).Implements(i3))

	AssertTrue(!reflect.TypeOf(t1).Implements(i2))
	AssertTrue(!reflect.TypeOf(t2).Implements(i1))
	AssertTrue(reflect.TypeOf(t3).Implements(i1))
	AssertTrue(reflect.TypeOf(t3).Implements(i2))

	var x Interface1 = Type1{}
	AssertEqual(x.func1(0), "type1")
	var y Interface2 = Type2{}
	AssertEqual(y.func2(""), 1)
	var z Interface3 = Type3{}
	AssertEqual(z.func1(0), "type1")
	AssertEqual(z.func2(""), 1)
	x = Type3{}
	AssertEqual(x.func1(0), "type1")
	y = Type2{}
	AssertEqual(y.func2(""), 1)

}

func checkEmptyInterface(){
	var any interface{}
	any = true
	AssertEqual(any, true)
	any = 12.34
	AssertEqual(any, 12.34)
	any = "hello"
	AssertEqual(any, "hello")
	any = map[string]int{"one": 1}
	AssertTrue(any != nil)
	any = new(bytes.Buffer)
	AssertTrue(any != nil)
}

func checkTypeAssertion(){
	var w io.Writer = os.Stdout
	var ok bool
	_, ok = w.(io.ReadWriter) // interface
	AssertTrue(ok)
	_, ok = w.(*os.File) // concrete type
	AssertTrue(ok)
	_, ok = w.(*bytes.Buffer) // concrete type
	AssertTrue(!ok)
	var t1 Interface1
	t1 = Type1{}
	_, ok = t1.(Interface2)
	AssertTrue(!ok)
	_, ok = t1.(Interface3)
	AssertTrue(!ok)
	var t2 Interface2
	t2 = Type2{}
	_, ok = t2.(Interface1)
	AssertTrue(!ok)
	_, ok = t2.(Interface3)
	AssertTrue(!ok)
	var t3 Interface3
	t3 = Type3{}
	_, ok = t3.(Interface1)
	AssertTrue(ok)
	_, ok = t3.(Interface2)
	AssertTrue(ok)
}

func checkTypeSwitch(){
	func1 := func(x interface{}) string{
		switch x.(type) {
		case int:
			return "int"
		case string:
			return "string"
		case bool:
			return "bool"
		default:
			return "unknown"
		}
	}
	AssertEqual(func1(1), "int")
	AssertEqual(func1(""), "string")
	AssertEqual(func1(true), "bool")
	AssertEqual(func1([]rune("")), "unknown")
}

func InterfaceMain(){
	checkIOWriteInterface()
	checkInterfaceImplementations()
	checkEmptyInterface()
	checkTypeAssertion()
}
