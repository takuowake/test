package main

import "fmt"

//You can declare "var" outside the function.
var (
	i int = 1
	f64 float64 = 1.2
	s string = "test"
	t, f bool = true, false
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
		Pi = 3.14
		Username = "test_user"
		Password = "test_pass"
	)
	var maximum_var int = 9223372036854775807
	const maximum_const = 9223372036854775807 + 1
	fmt.Println(Pi, Username, Password, maximum_var, maximum_const-1)
}

func main() {
	Var()
	foo()
	Const()
}
