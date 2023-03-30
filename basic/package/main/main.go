package main

import (
	"fmt"

	// 这里的github.com/wuvikr/go_study是go.mod中的module名
	// 即 go init 时所起的项目名称，后续basic/package/calu为包的路径名
	"github.com/wuvikr/go_study/basic/package/calc"
)

func main() {
	a := calc.Add(4, 5)
	fmt.Println(a)
}
