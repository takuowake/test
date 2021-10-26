package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func IncrementGenerator() (func() int) {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func Closure() {
	//x := 0
	//increment := func() int {
	//	x++
	//	return x
	//}
	//fmt.Println(increment())
	//fmt.Println(increment())
	//fmt.Println(increment())

	counter := IncrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println(c1(3))

	c2 := circleArea(3)
	fmt.Println(c2(2))
	fmt.Println(c2(3))
}

func Foo(params...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func VariadicArgument() {
	Foo()
	Foo(20, 30)
	Foo(20, 30, 40)

	s := []int{1, 2, 3}
	fmt.Println(s)
	Foo(s...)
}

func Exercise() {
	//Exercise1
	var f float64 = 1.11
	var i int = int(f)
	fmt.Println(i)

	f2 := 1.11
	i2 := int(f2)
	fmt.Println(i2)

	//Exercise3
	m := map[string]int{
		"Mike":20,
		"Nancy":24,
		"Messi":30,
	}
	fmt.Printf("%T %v", m, m)
}

func By2(num int)string {
	if num % 2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func If() {
	num := 100
	if num % 2 == 0 {
		fmt.Println("by 2")
	} else if num % 3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 20, 20
	if x == 20 && y == 20 {
		fmt.Println("&&")
	}
	if x == 20 || y == 20 {
		fmt.Println("||")
	}

	result := By2(10)
	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result)

	if result2 := By2(20); result2 == "ok" {
		fmt.Println("great2")
	}
}

func For() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	for {
		fmt.Println("hello")
	}
}

func Range() {
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	for i, v := range l {
		fmt.Println(i, v)
	}
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple":200, "banana":100}
	for k, i := range m {
		fmt.Println(k, i)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}
}

func GetOsName() string {
	return "ma"
}

func Switch() {
	switch os := GetOsName(); os {
	case "mac":
		fmt.Println("mac!!")
	case "windows":
		fmt.Println("windows!!")
	default:
		fmt.Println("default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	case t.Hour() > 17:
		fmt.Println("night")
	}
}

func Foo2() {
	fmt.Println("Hello foo2")
	fmt.Println("World foo2")
}
func Defer(){
	file, _ := os.Open("./var.go")
	/*
		Foo2()
		defer fmt.Println("World")
		fmt.Println("Hello")

		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

func LoggingSetting(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func Log() {
	LoggingSetting("test.log")
	_, err := os.Open("sakkaj")
	if err != nil {
		log.Fatalln("exit", err)
	}
	log.Println("test")
	log.Printf("%T %v", "test", "test")

	//log.Fatalln("error")
	log.Fatalf("%T %v", "test", "test")
}

func ErrorHandling() {
	file, err := os.Open("./var.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error!!")
	}
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
	//Func()
	//Closure()
	//VariadicArgument()
	//Exercise()
	//If()
	//For()
	//Range()
	//Switch()
	//Defer()
	//Log()
	ErrorHandling()
}

