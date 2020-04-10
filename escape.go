package main

type Addifier interface{ Add(a, b int32) int32 }

type Adder1 struct{ name string }

//go:noinline
func (adder Adder1) Add(a, b int32) int32 { return a + b }

func main() {
	adder := Adder1{name: "myadder"}
	adder.Add(10, 32)           //doesn't escape
	Addifier(adder).Add(10, 32) //escape
}
