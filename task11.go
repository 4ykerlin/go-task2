package main

import "fmt"

type Product struct {
	Name     string
	Category string
	Price    float64
}

func filterProducts(tovary []Product, maxCena float64, kategoriya string) []Product {
	var rezultat []Product
	// Оставляем товары дешевле maxCena и с нужной категорией
	for _, tovar := range tovary {
		if tovar.Price < maxCena && tovar.Category == kategoriya {
			rezultat = append(rezultat, tovar)
		}
	}
	return rezultat
}

func main() {
	tovary := []Product{
		{"Noski", "odezhda", 500},
		{"Telefon", "tehnika", 30000},
		{"Futbolka", "odezhda", 1500},
		{"Dzhinsy", "odezhda", 3500},
	}
	fmt.Println(filterProducts(tovary, 2000, "odezhda"))
}