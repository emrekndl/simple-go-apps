package foo

// Consumer interface and Dependency Injection
type bar interface {
	Merhaba()
}

type Foo struct {
	bar bar
}

func NewFoo(b bar) *Foo {
	return &Foo{
		bar: b,
	}
}

func (f *Foo) Greet() {
	f.bar.Merhaba()
}
