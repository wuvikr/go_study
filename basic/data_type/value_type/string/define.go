package main

import "fmt"

func main() {
	var a int
	var b string
	var c bool
	var d [5]int
	var e []string

	var s string = `         ,_---~~~~~----._
	_,,_,*^____      _____*g*\"*,--,
	/ __/ /'     ^.  /      \ ^@q   f
	[  @f | @))    |  | @))   l  0 _/
	\/   \~____ / __ \_____/     \
	|           _l__l_           I
	}          [______]           I
	]            | | |            |
	]             ~ ~             |
	|                            |
	 |                           |`

	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)
	fmt.Printf("d: %T\n", d)
	fmt.Printf("e: %T\n", e)

	fmt.Println(s)
}
