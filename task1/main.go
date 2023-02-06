package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от
	родительской структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int32
}

func (h Human) SayName() string {
	return h.name
}

func (h Human) SayAge() int32 {
	return h.age
}

// Структура Action, с ассоциацией Human
type Action struct {
	name string
	Human
}

func (a Action) SayName() string {
	return a.name
}

func main() {
	a := Action{name: "Ban", Human: Human{name: "Bill", age: 22}}
	fmt.Println(a.SayName())       // Out: Acion	(From Action)
	fmt.Println(a.Human.SayName()) // Out: Ban		(From Human)
	fmt.Println(a.SayAge())        // Out: 22		(From Human)
}
