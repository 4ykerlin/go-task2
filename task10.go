package main

import (
	"fmt"
	"strings"
)

type TextStats struct {
	Simvoly      int
	Slova        int
	Predlozhenia int
}

func textStats(tekst string) TextStats {
	var statistika TextStats
	// Количество символов (рун)
	statistika.Simvoly = len([]rune(tekst))
	// Количество слов через strings.Fields
	statistika.Slova = len(strings.Fields(tekst))

	// Считаем предложения по знакам . ! ?
	for _, simvol := range tekst {
		if simvol == '.' || simvol == '!' || simvol == '?' {
			statistika.Predlozhenia++
		}
	}
	return statistika
}

func main() {
	fmt.Println(textStats("Privet! Kak dela? Vse horosho."))
}