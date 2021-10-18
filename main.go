package main

import (
	"fmt"
	"strconv"
	"strings"
)

//You can declare "var" outside the function.
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func Var() {
	fmt.Println(i, f64, s, t, f)
}

//You cannot declare ":=" outside the function.
func foo() {
	xi := 1
	xi = 2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

func Const() {
	const (
		Pi       = 3.14
		Username = "test_user"
		Password = "test_pass"
	)
	var maximum_var int = 9223372036854775807
	const maximum_const = 9223372036854775807 + 1
	fmt.Println(Pi, Username, Password, maximum_var, maximum_const-1)
}

func NumericType() {
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 1.2
		c64 complex64 = 3 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v\n", u8, u8)
}

func Calculation() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("10 - 1 = ", 10-1)
	fmt.Println("10 / 2 = ", 10/2)
	fmt.Println("10 / 3 = ", 10/3)
	fmt.Println("10.0 / 3 = ", 10.0/3)
	fmt.Println("10 / 3.0 = ", 10/3.0)
	fmt.Println("10 % 2 = ", 10%2)
	fmt.Println("10 % 3 = ", 10%3)

	x := 0
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	fmt.Println(1 << 0)
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)
	fmt.Println(1 << 4)
}

func String() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println(`Test
                           Test
    Test`)
	fmt.Println(`"`)
}

func Bool() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 2, t)
	fmt.Printf("%T %v %t\n", f, 40, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}

func TypeConversion() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14g"
	ss, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("%T %v\n", ss, ss)

	H := "Hello World"
	fmt.Println(H[0])
}

func main() {
	//Var()
	//foo()
	//Const()
	//NumericType()
	//Calculation()
	//String()
	//Bool()
	TypeConversion()
}
