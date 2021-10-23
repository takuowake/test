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
	fmt.Println(string(H[0]))
}

func Array() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)
	var c []int = []int{100, 200}
	c = append(c, 200)
	fmt.Println(c)
}

func Slice() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[:2])
	fmt.Println(n[3:])
	fmt.Println(n[1:5])

	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)

	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)
}

func SliceMake() {
	a := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%d\n", len(a), cap(a), a)
	b := make([]int, 5, 5)
	fmt.Printf("len=%d cap=%d value=%d\n", len(b), cap(b), b)
	b = make([]int, 0)
	fmt.Printf("len=%d cap=%d value=%d\n", len(b), cap(b), b)

	var n []int
	fmt.Printf("len=%d cap=%d value=%d\n", len(n), cap(n), n)

	//m := make([]int, 0, 5)
	m := make([]int, 5)
	for i := 0; i < 5; i++ {
		m = append(m, i)
		fmt.Println(m)
	}
	fmt.Println(m)
}

func Map() {
	m := map[string]int{"apple": 100, "banana": 300}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["apple"] = 200
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)
	fmt.Println(m["nothing"])
	v, ok := m["nothing"]
	fmt.Println(v, ok)
	v2, ok2 := m["apple"]
	fmt.Println(v2, ok2)

	var m2 = make(map[string]int)
	m2["now"] = 5000
	fmt.Println(m2)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}

func Byte() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}

func add(x, y int) (int, int) {
	return x + y, x-y
}

func cal(price, item int) (result int){
	result = price * item
	return result
}

func Func() {
	fmt.Println(add(34, 26))
	r3 := cal(300, 4)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}

func main() {
	//Var()
	//foo()
	//Const()
	//NumericType()
	//Calculation()
	//String()
	//Bool()
	//TypeConversion()
	//Array()
	//Slice()
	//SliceMake()
	//Map()
	//Byte()
	Func()
}
