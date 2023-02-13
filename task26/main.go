package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет, что все символы в
	строке уникальные (true — если уникальные, false etc). Функция
	проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func unique(s string) bool {
	//Переводим строку в нижний регистр и переводим в руны
	runes := []rune(strings.ToLower(s))

	//символы
	runesSet := make(map[rune]struct{})

	//Добовляем руны в карту
	for _, ch := range runes {
		//Если руна в карте то повтор
		if _, ok := runesSet[ch]; ok {
			return false
		}
		runesSet[ch] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(unique("abcd"))      //Out: true
	fmt.Println(unique("abCdefAaf")) //Out: false
	fmt.Println(unique("aabcd"))     //Out: false
	fmt.Println(unique("aAbcd"))     //Out: false
	fmt.Println(unique("a😌bcd"))     //Out: true
}
