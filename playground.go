package main

import "fmt"

type MyStruct struct {
	name string
}

func (ms MyStruct) MethodA() string {
	return ms.name
}

func (ms MyStruct) MethodB() string {
	return ms.name
}

func (ms MyStruct) MethodC() string {
	return ms.name
}

func (ms MyStruct) MethodD() string {
	return ms.name
}


func main() {
	newStruct := MyStruct{name: "MethodA"}

	result := newStruct.MethodD()

	fmt.Println("Result is " + result)

}