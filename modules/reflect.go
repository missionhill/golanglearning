package modules

import (
	"reflect"
	"io"
	"os"
	"time"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}


func CheckType(){
	AssertEqual(reflect.TypeOf(3).String(), "int")
	v1 := reflect.ValueOf(3)
	AssertEqual(v1.Type().String(), "int")
	AssertEqual((v1.Interface()).(int), 3)
	AssertEqual(v1.Kind(), reflect.Int)
	AssertEqual(reflect.TypeOf(3.1).String(), "float64")
	v2 := reflect.ValueOf(3.1)
	AssertEqual(v2.Type().String(), "float64")
	AssertEqual(v2.Kind(), reflect.Float64)
	var w io.Writer = os.Stdout
	AssertEqual(reflect.TypeOf(w).String(), "*os.File")
	v3 := reflect.ValueOf(os.Stdout)
	AssertEqual(v3.Type().String(), "*os.File")
	AssertEqual(v3.Kind(), reflect.Ptr)
	x := 1
	v4 := &x
	AssertEqual(reflect.TypeOf(v4).String(), "*int")
	AssertEqual(reflect.ValueOf(v4).Type().String(), "*int")
	AssertTrue(reflect.ValueOf(v4).Elem().Int()==1)
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	AssertTrue(reflect.ValueOf(strangelove).Type().Field(0).Name=="Title")
	v5 := reflect.ValueOf(&x).Elem()
	v5.Set(reflect.ValueOf(4))
	AssertTrue(x==4)
	v5.SetInt(5)
	AssertTrue(x==5)
	v6 := reflect.ValueOf(time.Hour).Type().String()
	AssertEqual(v6, "time.Duration")
	v7 := reflect.ValueOf(time.Hour)
	AssertEqual(v7.NumMethod(), 5)
}

func ReflectMain(){
	CheckType()
}
