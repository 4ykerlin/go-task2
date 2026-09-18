package main

import "fmt"

// Функция, которая считает голоса
func podschetGolosov(golosa []string) {
	// Создаем переменные-счетчики для каждого кандидата
	var anna int = 0
	var boris int = 0
	var viktor int = 0

	// Обычным циклом проходим по всему массиву с голосами
	for i := 0; i < len(golosa); i++ {
		if golosa[i] == "Анна" {
			anna = anna + 1
		}
		if golosa[i] == "Борис" {
			boris = boris + 1
		}
		if golosa[i] == "Виктор" {
			viktor = viktor + 1
		}
	}

	// Считаем общее число проголосовавших
	// Переводим во float64, чтобы правильно посчитались проценты с точкой
	var vsego float64 = float64(len(golosa))

	// Считаем проценты для каждого
	var procentAnna float64 = (float64(anna) / vsego) * 100
	var procentBoris float64 = (float64(boris) / vsego) * 100
	var procentViktor float64 = (float64(viktor) / vsego) * 100

	// Выводим результаты на экран
	fmt.Println("Результаты голосования:")
	fmt.Println("Анна:", anna, "голосов,", procentAnna, "%")
	fmt.Println("Борис:", boris, "голосов,", procentBoris, "%")
	fmt.Println("Виктор:", viktor, "голосов,", procentViktor, "%")
}

func main() {
	// 1. Создаем срез с именами кандидатов
	kandidati := []string{"Анна", "Борис", "Виктор"}
	fmt.Println("Список кандидатов:", kandidati)

	// Имитируем массив с голосами, которые поступили
	vseGolosovali := []string{"Анна", "Виктор", "Борис", "Анна", "Анна", "Виктор"}

	// 2. Вызываем нашу функцию
	podschetGolosov(vseGolosovali)
}
