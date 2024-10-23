package main

import (
	"interface-struct/internal/bar"
	"interface-struct/internal/foo"
)

func main() {
	bar := &bar.Bar{}
	foo := foo.NewFoo(bar)

	foo.Greet()
}
