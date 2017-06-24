package modules

type Demo struct{}

func (d Demo) func1() string {
	return "func1"
}
func (d *Demo) func2() string {
	return "func2"
}

type Child struct{
	X string
}
type Parent struct{Child} // embedded type Child

func (child Child) whoAmI() string{
	return "child"
}

type Int int

func(i Int) func3(val int){
	i = Int(val)
}

func(i *Int) func4(val int){
	*i = Int(val)
}

func checkDeclaration(){
	var d Demo
	AssertEqual(d.func1(), "func1")
	d = Demo{}
	AssertEqual(d.func1(), "func1")
	AssertEqual(Demo{}.func1(), "func1")
	AssertEqual((&d).func1(), "func1")
	AssertEqual((&d).func2(), "func2")
	AssertEqual(d.func2(), "func2")
	AssertEqual( (*(&d)).func2(), "func2")

	var i Int
	i.func3(3)
	AssertEqual(int(i),0)
	i.func4(4)
	AssertEqual(int(i), 4)

	// invoking via instances
	method1 := d.func1
	AssertEqual(method1(), "func1")
	method2 := (&d).func2
	AssertEqual(method2(), "func2")

	// invoking via type
	method3 := Demo.func1
	AssertEqual(method3(d), "func1")
	method4 := (*Demo).func2
	AssertEqual(method4(&d), "func2")
}

func checkEmbedded(){
	p := Parent{}
	AssertEqual(p.X, "")
	AssertEqual(p.whoAmI(), "child")
}

func MethodMain(){
	checkDeclaration()
	checkEmbedded()
}