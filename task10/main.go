package main

import (
	"fmt"
	"math"
)

/*
	Дана последовательность температурных колебаний: -25.4,
	-27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить
	данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0,
	15.5}, 20: {24.5}, etc.
*/

func temperatureGroups(temperatures []float64) map[int][]float64 {
	groupedTemperatures := make(map[int][]float64)

	for _, temperature := range temperatures {
		// Выносим знак из флора для правильного округления отрецательных значений
		group := int(math.Abs(temperature)/temperature*math.Floor(math.Abs(temperature/10.0))) * 10
		// Используем map для групировки
		groupedTemperatures[group] = append(groupedTemperatures[group], temperature)
	}

	return groupedTemperatures
}

func main() {
	teratures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	gropes := temperatureGroups(teratures)

	for key, value := range gropes {

		fmt.Printf("%d: [", key)
		for _, t := range value {
			fmt.Printf("%f, ", t)
		}
		fmt.Printf("]\n")
	}
}
