package main

import "fmt"

func main() {
	// создаем переменные для каждого веса
	var osnovnoibagah float64
	var rychnaclad float64
	var doprychnaclad float64

	// По очереди просим пользователя ввести числа
	fmt.Println("Введите вес большого чемодана:")
	fmt.Scan(&osnovnoibagah)

	fmt.Println("Введите вес ручной клади:")
	fmt.Scan(&rychnaclad)

	fmt.Println("Введите вес доп багажа:")
	fmt.Scan(&doprychnaclad)

	// Считаем сумму 
	var vsego float64
	vsego = osnovnoibagah + rychnaclad + doprychnaclad

	// Выводим
	fmt.Println("Общий вес багажа:", vsego , "кг")
}
