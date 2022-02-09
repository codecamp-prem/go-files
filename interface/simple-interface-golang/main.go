package main

import "fmt"

// MyInterface : a interface
type MyInterface interface {
	MethodWithoutParameter()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

// MyType : a defined type
type MyType int

// MethodWithoutParameter : a method for MyType
func (m MyType) MethodWithoutParameter() {
	fmt.Println("Calling from MethodWithoutParameter()")
}

// MethodWithParameter : a method for MyType
func (m MyType) MethodWithParameter(f float64) {
	fmt.Printf("Calling from MethodWithParameter() %.2f\n", f)
}

// MethodWithReturnValue : a method for MyType
func (m MyType) MethodWithReturnValue() string {
	return "Calling from MethodWithReturnValue()"
}

func main() {
	var value MyInterface
	value = MyType(5)
	value.MethodWithoutParameter()
	value.MethodWithParameter(99.9)
	fmt.Println(value.MethodWithReturnValue())
}
