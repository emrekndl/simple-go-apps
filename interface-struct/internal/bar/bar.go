package bar

import (
	"fmt"
)

type Bar struct{}

func (b Bar) SayHello() {
	fmt.Println("Hello!")
}

func (b Bar) SayHola() {
	fmt.Println("Hola!")
}

func (b Bar) SayPrivet() {
	fmt.Println("Privet!")
}

func (b Bar) Merhaba() {
	fmt.Println("Merhaba!")
}
