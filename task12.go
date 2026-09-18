package main

import (
	"fmt"
	"strconv"
)

// Константы для систем счисления
const (
	bin = 2
	dec = 10
	hex = 16
)

// Универсальная функция конвертации
func convert(chisloStr string, izBase, vBase int) (string, error) {
	chislo, oshibka := strconv.ParseInt(chisloStr, izBase, 64)
	if oshibka != nil {
		return "", oshibka
	}
	return strconv.FormatInt(chislo, vBase), nil
}

func main() {
	var chislo int64
	fmt.Scan(&chislo)

	// Выводим в двоичной, десятичной и шестнадцатеричной
	fmt.Println("Bin:", strconv.FormatInt(chislo, bin))
	fmt.Println("Dec:", strconv.FormatInt(chislo, dec))
	fmt.Println("Hex:", strconv.FormatInt(chislo, hex))

	// Пример конвертации из hex в bin
	rezultat, _ := convert("ff", hex, bin)
	fmt.Println("ff ->", rezultat)
}