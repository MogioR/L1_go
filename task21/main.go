package main

import (
	"fmt"
	"math"
)

/*
	Реализовать паттерн «адаптер» на любом примере.
*/

// Интерфейс колеса
type Wheel interface {
	GetWheelHeight() int32
	SetWheelHeight(a int32)
}

// Круглое колесо часть системы, подходит под интерфейс
type CircleWheel struct {
	radius int32
}

func (c CircleWheel) GetWheelHeight() int32 {
	return c.radius
}

func (c *CircleWheel) SetWheelHeight(a int32) {
	c.radius = a
}

// Квадратное колесо (не подходит под интерфейс и нельзя менять)
type Square struct {
	side int32
}

func (s Square) GetSide() int32 {
	return s.side
}

func (c *Square) SetSide(a int32) {
	c.side = a
}

// Создаём адаптер
type SquareWheel struct {
	Square
}

// Добовляем методы
func (s SquareWheel) GetWheelHeight() int32 {
	return int32(int(float64(s.GetSide()) / math.Sqrt2))
}

func (c *SquareWheel) SetWheelHeight(a int32) {
	c.SetSide(int32(int(math.Ceil(float64(a) * math.Sqrt2))))
}

func main() {
	// Слайс на 2 колеса
	wheels := make([]Wheel, 2)
	// Круглое и квадратное
	wheels[0] = &CircleWheel{50}
	wheels[1] = &SquareWheel{Square: Square{50}}

	fmt.Printf(
		" Circle wheel height: %d\n Square wheel height: %d\n",
		wheels[0].GetWheelHeight(),
		wheels[1].GetWheelHeight(),
	)

	wheels[0].SetWheelHeight(10)
	wheels[1].SetWheelHeight(10)

	fmt.Printf(
		" Circle wheel height: %d\n Square wheel height: %d\n",
		wheels[0].GetWheelHeight(),
		wheels[1].GetWheelHeight(),
	)
}
