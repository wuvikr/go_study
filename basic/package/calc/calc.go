package calc

// 在go中，函数，变量首字母是否大小写决定了其是否为公开的还是私有的
// 首字母大写表示公开变量，可以被其他包访问使用
// 首字母小写表示私有变量，
var Name = "zhangsan"
var Age = "19"

func Add(x, y int) int {
	return x + y
}
