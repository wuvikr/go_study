package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("hello, %s\n", name)
}

func f1(name string, f func(string)) {
	f(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc1(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func calc2(op string) func(int, int) int {
	switch op {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	f1("Alice", sayHello)

	add_result := calc1(10, 20, add)
	fmt.Printf("add_result: %v\n", add_result)

	sub_result := calc1(30, 10, sub)
	fmt.Printf("sub_result: %v\n", sub_result)

	fAdd := calc2("+")
	fmt.Printf("fAdd(10, 20): %v\n", fAdd(10, 20))

	fSub := calc2("-")
	fmt.Printf("fSub(30, 10): %v\n", fSub(30, 10))

}
