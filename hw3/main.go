package main

import (
	"fmt"
)

func main() {

	type car struct {
		model     string // марка
		release   int    // год выпуска
		trunk     int    // обьем багажника
		engine    string // запущен ли двигатель
		window    string // открыли ли окна
		trunkload int    // обьем загрузки багажника
	}

	type truck struct {
		model     string
		release   int
		trunk     int
		engine    string
		window    string
		trunkload int
	}

	var a car
	var b truck
	a = car{"C5", 2019, 34, "yes", "no", 22}
	b = truck{model: "T10", release: 1987, trunk: 250, engine: "yes", window: "yes", trunkload: 135}
	fmt.Println(a)
	fmt.Println(b)
	b = truck{model: "T10", engine: "yes", window: "yes"}
	fmt.Println(b)
	fmt.Printf("Model: %+v\nRelease: %+v\nTrunk: %+v\nEngine: %+v\nWindow: %+v\nTrunkload: %+v\n", a.model, a.release, a.trunk, a.engine, a.window, a.trunkload)

	/*	Вывод:
		{C5 2019 34 yes no 22}
		{T10 1987 250 yes yes 135}
		{T10 0 0 yes yes 0}
		Model: C5
		Release: 2019
		Trunk: 34
		Engine: yes
		Window: no
		Trunkload: 22
	*/
	fmt.Println()
	var a1 = car{"C5", 2019, 34, "yes", "no", 22}
	var a2 = a1
	var a3 = car{"WW338", 2019, 21, "no", "no", 14}
	var b2 = truck{model: "T10", release: 1987, trunk: 250, engine: "yes", window: "yes", trunkload: 135}

	// Сравнение полей
	if a1.release == a2.release && a1.release == a3.release {
		fmt.Println("Год выпуска a1, a2 и а3 одинаковый ")
	}
	fmt.Println(a1.trunkload == b2.trunkload)
	/*Вывод:
	Год выпуска a1, a2 и а3 одинаковый
	false
	*/
	fmt.Println()
	// Сравнение структур
	if a1 == a2 {
		if a1 != a3 {
			fmt.Println("a1 равен а2 и не равен а3")
		} else {
			fmt.Println("a1 равен а2 и а3")
		}
	} else if a1 != a2 {
		if a1 == a3 {
			fmt.Println("a1 не равен a2, но равен а3")
		} else {
			fmt.Println("a1 не равен а2 и а3")
		}
	}
	// a1 равен а2 и не равен а3
	fmt.Println()

	// Очередь стек

	push("Я первый в очереди")
	push("Я второй в очереди")
	push("Я третий в очереди")
	push("Я последний в очереди")
	push("Я следующий в очереди 1")
	fmt.Println(pop())
	fmt.Println(pop())
	push("Я следующий в очереди 2")
	fmt.Println(pop())
	fmt.Println(pop())
	push("Я следующий в очереди 3")
	push("Я в какой то в очереди")
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
}

var x []string // переменная для стека

func push(str string) {
	x = append(x, str)
}

func pop() string {
	numOfElements := len(x)
	if numOfElements == 0 {
		return "Никого нет"
	}
	popElem := x[0]
	x = x[1:numOfElements]
	return popElem
}

/*
Я первый в очереди
Я второй в очереди
Я третий в очереди
Я последний в очереди
Я следующий в очереди 1
Я следующий в очереди 2
Я следующий в очереди 3
Я в какой то в очереди
Никого нет
*/
