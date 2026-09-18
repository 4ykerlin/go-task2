package main

import "fmt"

type InventoryItem struct {
	Name        string
	Weight      float64
	IsQuestItem bool
}

func obshiyVes(predmety []InventoryItem) float64 {
	summa := 0.0
	// Складываем вес всех предметов
	for _, predmet := range predmety {
		summa += predmet.Weight
	}
	return summa
}

func main() {
	inventar := []InventoryItem{
		{"Mech", 5.5, false},
		{"Shchit", 7.0, false},
		{"Zelye", 0.5, true},
		{"Kluch", 0.2, true},
		{"Bronya", 12.0, false},
	}
	fmt.Println("Obshiy ves:", obshiyVes(inventar))
}