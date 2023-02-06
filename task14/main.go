package main

import (
	"fmt"
	"reflect"
)

/*
	Разработать программу, которая в рантайме способна определить тип
	переменной: int, string, bool, channel из переменной типа interface{}.
*/

func getTypeSwich(a interface{}) string {
	// Проверяем возможность приведения типа
	switch a.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		return "unknown"
	}
}

func getTypeIf(a interface{}) string {
	// Проверяем возможность приведения типа
	if _, ok := a.(int); ok {
		return "int"
	} else if _, ok := a.(string); ok {
		return "string"
	} else if _, ok := a.(bool); ok {
		return "bool"
	} else if _, ok := a.(chan int); ok {
		return "chan int"
	} else if _, ok := a.(chan string); ok {
		return "chan string"
	} else if _, ok := a.(chan bool); ok {
		return "chan bool"
	} else {
		return "unknown"
	}
}

func getTypeReflect(a interface{}) string {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "unknown"
	}
}

func main() {
	//If style cast
	fmt.Printf("%s\n", getTypeIf(1))              // Out: int
	fmt.Printf("%s\n", getTypeIf("dfsfds"))       // Out: string
	fmt.Printf("%s\n", getTypeIf(true))           // Out: bool
	fmt.Printf("%s\n", getTypeIf(make(chan int))) // Out: chan int
	fmt.Printf("%s\n", getTypeIf(1.3))            // Out: unknown

	fmt.Printf("\n")

	//Swich style cast
	fmt.Printf("%s\n", getTypeSwich(1))              // Out: int
	fmt.Printf("%s\n", getTypeSwich("dfsfds"))       // Out: string
	fmt.Printf("%s\n", getTypeSwich(true))           // Out: bool
	fmt.Printf("%s\n", getTypeSwich(make(chan int))) // Out: chan int
	fmt.Printf("%s\n", getTypeSwich(1.3))            // Out: unknown

	fmt.Printf("\n")

	//Reflect method
	fmt.Printf("%s\n", getTypeReflect(1))              // Out: int
	fmt.Printf("%s\n", getTypeReflect("dfsfds"))       // Out: string
	fmt.Printf("%s\n", getTypeReflect(true))           // Out: bool
	fmt.Printf("%s\n", getTypeReflect(make(chan int))) // Out: chan int
	fmt.Printf("%s\n", getTypeReflect(1.3))            // Out: unknown
}
