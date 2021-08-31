package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// var buffer bytes.Buffer
	// buffer.WriteString("ddddddd")
	// buffer.WriteString(111)
	s := "ddddddd%v"
	var i int64 = 1111
	l := s + strconv.FormatInt(i, 10)

	y := fmt.Sprintf(s, i)
	fmt.Println("=====================")
	fmt.Println(y)
	fmt.Println("=====================")

	fmt.Println("s = ", l)

	ss := "你好，我是邹超ABC我会，xyz我也会"

	b := strings.Contains(strings.ToLower(ss), strings.ToLower("我是邹超"))
	fmt.Printf("%#v\n", b)

}
