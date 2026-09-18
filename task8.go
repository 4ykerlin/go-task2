package main

import "fmt"

// 1. Структура для лога
type LogZapis struct {
	IPaddress string
	KodOshibki int
	Vremya     string
}

// 3. Функция ищет только коды 4xx и 5xx
func najdiOshibki(vseLogi []LogZapis) []LogZapis {
	var tolkoOshibki []LogZapis
	
	for i := 0; i < len(vseLogi); i++ {
		stroka := vseLogi[i]
		// Проверяем код от 400 до 599
		if stroka.KodOshibki >= 400 && stroka.KodOshibki <= 599 {
			tolkoOshibki = append(tolkoOshibki, stroka)
		}
	}
	
	return tolkoOshibki
}

func main() {
	// 2. Срез логов
	moiLogi := []LogZapis{
		{IPaddress: "192.168.1.1", KodOshibki: 200, Vremya: "12:00"},
		{IPaddress: "192.168.1.2", KodOshibki: 404, Vremya: "12:05"},
		{IPaddress: "192.168.1.3", KodOshibki: 500, Vremya: "12:10"},
	}
	
	plohieLogi := najdiOshibki(moiLogi)
	fmt.Println(plohieLogi)
}
