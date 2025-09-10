package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var a, b big.Int
	fmt.Println("Введите два числа")
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)
}

func Add(a, b big.Int) big.Int {
	z := big.NewInt(0)
	z.Add(&a, &b)
	return *z
}

func Subtract(a, b big.Int) big.Int {
	z := big.NewInt(0)
	z.Sub(&a, &b)
	return *z
}

func Multiply(a, b big.Int) big.Int {
	z := big.NewInt(0)
	z.Mul(&a, &b)
	return *z
}

func Division(a, b big.Int) big.Int {
	z := big.NewInt(0)
	z.Div(&a, &b)
	return *z
}
