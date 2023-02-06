package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными
	параметрами x,y и конструктором.
*/

type Point struct {
	x int32
	y int32
}

// New для создания указателя
func NewPoint(x, y int32) *Point {
	return &Point{x, y}
}

// Make для создания экземпляра
func MakePoint(x, y int32) Point {
	return Point{x, y}
}

func (p *Point) SetPos(x, y int32) {
	p.x = x
	p.y = y
}

func (p Point) GetPos() (x, y int32) {
	return p.x, p.y
}

func (p Point) GetDistanseTo(a Point) float64 {
	return math.Sqrt(math.Pow(float64(p.x-a.x), 2) + math.Pow(float64(p.y-a.y), 2))
}

func main() {
	a := NewPoint(10, 10)
	b := NewPoint(5, 5)

	fmt.Println(a.GetDistanseTo(*b))

}
