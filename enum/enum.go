package main

import "fmt"

type MyEnum int

const (
	Foo MyEnum = 1
	Bar MyEnum = 2
)

func (e MyEnum) String() string {
	switch e {
	case Foo:
		return "Foo"
	case Bar:
		return "Bar"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type MyStruct struct {
	Field1 MyEnum
	field2 MyEnum
}

func main() {
	info := &MyStruct{
		Field1: MyEnum(Foo),
		field2: MyEnum(Bar),
	}
	fmt.Printf("%v\n", info)
	fmt.Printf("%+v\n", info)
	fmt.Printf("%#v\n", info)
}
