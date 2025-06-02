package main

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func _pointers() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func _stringsAndRunes() {
	const s = "สวัสดี"

	// NOTE: The length of a string is the number of bytes. In Go the strings is a read-only slice of bytes.
	fmt.Println("len: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune Count: ", utf8.RuneCountInString(s)) // 6

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 24
	return &p
}

func _structs() {
	fmt.Println(person{"Bob", 22})
	fmt.Println(person{name: "Alice", age: 23})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Umay", age: 27})

	fmt.Println(newPerson("Ann"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{name: "Fox", isGood: true}
	fmt.Println(dog)
}

type _rect struct {
	width, height int
}

func (r *_rect) area() int {
	// r.width = 10
	return r.width * r.height
}

func (r _rect) perim() int {
	// r.width = 10
	return 2*r.width + 2*r.height
}

func _methods() {
	r := _rect{width: 19, height: 27}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	// fmt.Println("rp: ", rp)
	// fmt.Println("r: ", r)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func _interfaces() {
	r := rect{
		width:  7,
		height: 9,
	}

	c := circle{
		radius: 13,
	}

	measure(r)
	measure(c)
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
	return StateConnected
}

func _enums() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func _structEmbedding() {
	co := container{
		base: base{num: 29},
		str:  "go",
	}
	fmt.Printf("co=(num: %v, str: %v)\n", co.num, co.str)
	fmt.Printf("co.base.num: %v\n", co.base.num)
	fmt.Printf("co.describe(): %v\n", co.describe())

	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer", d.describe())
}

func mapKeys[K comparable, V any](m map[K]V) []K {
	// mapKeys is a copy of the map's keys in a slice.
	// NOTE: K is comparable value and V is any type. [K comparable, V any] this block is a type declaration for a generic function.
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func (lst *List[T]) Pop() (T, bool) {
	if lst.head == nil {
		var zero T
		return zero, false
	}
	val := lst.head.val
	lst.head = lst.head.next
	if lst.head == nil {
		lst.tail = nil
	}
	return val, true

}

/* func (lst *List[T]) Pop() (T, bool) {
	if lst.tail == nil {
		var zero T
		return zero, false
	}
	val := lst.tail.val
	if lst.head == lst.tail {
		lst.head = nil
		lst.tail = nil
	} else {
		prev := lst.head
		for prev.next != lst.tail {
			prev = prev.next
		}
		prev.next = nil
		lst.tail = prev
	}
	return val, true
} */

func _generics() {
	var m = map[int]string{
		1: "0",
		3: "2",
		5: "4",
	}
	fmt.Println("Map Keys:", mapKeys(m))
	_ = mapKeys[int, string](m)

	l := List[int]{}
	l.Push(17)
	l.Push(29)
	l.Push(97)
	fmt.Println("List:", l.GetAll())

	l2 := List[string]{}
	l2.Push("foo")
	l2.Push("bar")
	l2.Push("baz")
	fmt.Println("List:", l2.GetAll())

	v, ok := l.Pop()
	fmt.Println("Pop:", v, ok)
	v, ok = l.Pop()
	fmt.Println("Pop:", v, ok)
	v, ok = l.Pop()
	fmt.Println("Pop:", v, ok)
	v, ok = l.Pop()
	fmt.Println("Pop:", v, ok)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrorOutOfTea = fmt.Errorf("no more tea avilable")
var ErrorPower = fmt.Errorf("cant't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("Making tea: %w", ErrorPower)
	}
	return nil
}

func _errors() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrorOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrorPower) {
				fmt.Println("Now it's dark.")
			} else {
				fmt.Printf("Unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func _customErrors() {
	_, err := f2(42)
	var arge *argError
	if errors.As(err, &arge) {
		fmt.Println("Arg error:", arge.arg)
		fmt.Println("Message:", arge.message)
	} else {
		fmt.Println("Error does not matching argError")
	}
}

func main() {
	// NOTE: Examples 15..24 are from https://gobyexample.com/
	// Pointers
	_pointers()
	fmt.Println("---------------")
	// Strings and Runes
	_stringsAndRunes()
	fmt.Println("---------------")
	// Structs
	_structs()
	fmt.Println("---------------")
	// Methods
	_methods()
	fmt.Println("---------------")
	// Interfaces
	_interfaces()
	fmt.Println("---------------")
	// Enums
	_enums()
	fmt.Println("---------------")
	// Struct Embedding
	_structEmbedding()
	fmt.Println("---------------")
	// Generics
	_generics()
	fmt.Println("---------------")
	// Errors
	_errors()
	fmt.Println("---------------")
	// Custom Errors
	_customErrors()
	fmt.Println("---------------")
}
