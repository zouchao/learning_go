package main

import "fmt"

func main() {
	// 1. 声明一个变量，默认值为0
	var a int
	fmt.Println("a = ", a)
	// 2. 声明一个变量，初始化值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("b type of = %T\n", b)
	// 3 初始化是，可以省去数据类型，通过值自动匹配当前变量的数据类型
	var bb = "abcd"
	fmt.Printf("bb = %s, type of %T\n", bb, bb)
	cc := 100
	fmt.Println("cc = ", cc)
	fmt.Printf("%T\n", cc)
	fmt.Printf("%d\n", cc)
}
