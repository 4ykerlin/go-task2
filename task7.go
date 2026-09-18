package main

import "fmt"

// 1. Структура для сотрудника
type Sotrudnik struct {
	ID       int
	Imya     string
	Dolzhnost string
	Zarplata float64
}

// 3. Функция считает общую и среднюю зарплату
func schitajDengi(vseSotrudniki []Sotrudnik) (float64, float64) {
	var obshiyFond float64 = 0
	
	for i := 0; i < len(vseSotrudniki); i++ {
		obshiyFond = obshiyFond + vseSotrudniki[i].Zarplata
	}
	
	// Считаем среднюю зарплату
	var kolvo float64 = float64(len(vseSotrudniki))
	var srednyaya float64 = obshiyFond / kolvo
	
	return obshiyFond, srednyaya
}

func main() {
	// 2. Срез с данными
	lyudi := []Sotrudnik{
		{ID: 1, Imya: "Ivan", Dolzhnost: "Programmist", Zarplata: 150000},
		{ID: 2, Imya: "Anna", Dolzhnost: "Dizajner", Zarplata: 90000},
	}
	
	fond, srednee := schitajDengi(lyudi)
	
	fmt.Println("Всего денег:", fond)
	fmt.Println("Средняя:", srednee)
}
