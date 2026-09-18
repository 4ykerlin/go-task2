package main

import "fmt"

// Функция собирает уникальные теги
func soberiTegi(vsePosti [][]string) []string {
	// Карта для проверки уникальности
	proverka := make(map[string]bool)
	
	for i := 0; i < len(vsePosti); i++ {
		odinPost := vsePosti[i]
		for j := 0; j < len(odinPost); j++ {
			teg := odinPost[j]
			proverka[teg] = true // Записываем, что тег существует
		}
	}
	
	// Сюда складываем результат
	var rezultat []string
	for klyuch := range proverka {
		rezultat = append(rezultat, klyuch)
	}
	
	return rezultat
}

func main() {
	// 1. Создаем срез срезов
	moiPosti := [][]string{
		{"go", "backend"},
		{"git", "go", "tools"},
	}
	
	// 2. Вызываем функцию
	itog := soberiTegi(moiPosti)
	
	// 3. Выводим результат
	fmt.Println(itog)
}

