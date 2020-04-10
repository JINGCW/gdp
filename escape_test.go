package main

import "testing"

func BenchmarkDirect(b *testing.B) {
	adder := Adder1{name: "6754"}
	for i := 0; i < b.N; i++ {
		adder.Add(10, 32)
	}
}

func BenchmarkInterface(b *testing.B) {
	adder := Adder1{name: "6754"}
	for i := 0; i < b.N; i++ {
		Addifier(adder).Add(10, 32)
	}
}
