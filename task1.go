package main

import "fmt"

func main() {
	// Стоимость проживания
	const rabochiedni = 2100
	const vihodniedni = 2850

	// Даты с галочками на бланке (Сентябрь 2025):
	// Будние (ВТ 9, СР 10, ЧТ 11, ЧТ 25) — 4 дня
	// Выходные (ПТ 12, СБ 13, ВС 14, ПТ 26) — 4 дня
	rabochiedniCount := 4
	vihodniedniCount := 4

	totalCost := (rabochiedniCount * rabochiedni) + (vihodniedniCount * vihodniedni)

	fmt.Println(totalCost)
}
