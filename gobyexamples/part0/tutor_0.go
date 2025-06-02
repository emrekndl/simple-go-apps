package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"time"
)

const s string = "constant"

func _variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

func _constants() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}

func _forloop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for x := range 3 {
		fmt.Println("range: ", x)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

func _ifelse() {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative.")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit.")
	} else {
		fmt.Println(num, "has multiple digits.")
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

}

func _switch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("integer")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("heyho")
}

func _arrays() {
	var a [5]int
	fmt.Println("empty: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

func _slices() {
	var s []string
	fmt.Println("uninitialized: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:5]
	fmt.Println("slice[x:y]: ", l)

	l = s[:5]
	fmt.Println("slice[:y]: ", l)

	l = s[2:]
	fmt.Println("slice[x:]: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declaration: ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func _maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	clear(m)
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("if exists key: ", prs)

	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map literal: ", n)

	n2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

func _range() {
	nums := []int{1, 2, 3, 8}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 8 {
			fmt.Println("index: ", i)
		}
	}

	string_map := map[string]string{
		"key1": "foo",
		"key2": "bar",
	}

	for k, v := range string_map {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range string_map {
		fmt.Println("key: ", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}

func _functions() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

func vals() (int, int) {
	return 17, 29
}

func _multipleReturnValues() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

func sum(nums ...int) {
	fmt.Println("nums: ", nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println("sum: ", total)
}

func _variadicFunctions() {
	sum(1, 2)
	sum(1, 2, 7)

	nums := []int{2, 3, 5, 7}
	sum(nums...)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func _closures() {
	// NOTE: Closures are anonymous functions with access to variables that exist outside of their body.
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print("fact: ", n, " ")
	return n * fact(n-1)
}

func _recursion() {
	fmt.Println(fact(5))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		fmt.Print("fib: ", n, " ")
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(5))
}

func main() {
	// NOTE: Examples 1..14 are from https://gobyexample.com/
	//  Variables
	_variables()
	fmt.Println("---------------")
	// Constants
	_constants()
	fmt.Println("---------------")
	// For Loop
	_forloop()
	fmt.Println("---------------")
	// If Else
	_ifelse()
	fmt.Println("---------------")
	// Switch
	_switch()
	fmt.Println("---------------")
	//Arrays
	_arrays()
	fmt.Println("---------------")
	// Slices
	_slices()
	fmt.Println("---------------")
	// Maps
	_maps()
	fmt.Println("---------------")
	// Range
	_range()
	fmt.Println("---------------")
	// Functions
	_functions()
	fmt.Println("---------------")
	// Multiple Return Values
	_multipleReturnValues()
	fmt.Println("---------------")
	// Variadic Functions
	_variadicFunctions()
	fmt.Println("---------------")
	// Closures
	_closures()
	fmt.Println("---------------")
	// Recursion
	_recursion()
	fmt.Println("---------------")

}
