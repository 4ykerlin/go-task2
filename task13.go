package main

import "fmt"

func main() {
	// Карта трат по категориям
	trakom := map[string]float64{
		"Eda":           15000,
		"Transport":     5000,
		"Razvlecheniya": 3000,
	}

	// Добавляем новые траты в категорию "Еда"
	trakom["Eda"] += 2000

	// Считаем итоговую сумму по всем категориям
	itog := 0.0
	for _, summa := range trakom {
		itog += summa
	}
	fmt.Println("Itogovaya summa:", itog)
}