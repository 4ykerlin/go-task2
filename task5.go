package main

import (
	"errors"
	"fmt"
	"strings"
)

// Функция проверки, как в задании
func validateUser(name string, age int, email string) error {
	if name == "" {
		return errors.New("ошибка: имя не должно быть пустым")
	}
	if len(name) >= 50 {
		return errors.New("ошибка: имя должно быть меньше 50 символов")
	}
	if age < 18 || age > 120 {
		return errors.New("ошибка: возраст должен быть от 18 до 120 лет")
	}
	if strings.Index(email, "@") == -1 {
		return errors.New("ошибка: email должен содержать символ @")
	}
	return nil
}

func main() {
	// Создаем переменные для ввода
	var imya string
	var vozrast int
	var pochta string

	// Спрашиваем имя
	fmt.Print("Введите имя: ")
	fmt.Scan(&imya)

	// Спрашиваем возраст
	fmt.Print("Введите возраст: ")
	fmt.Scan(&vozrast)

	// Спрашиваем email
	fmt.Print("Введите email: ")
	fmt.Scan(&pochta)

	// Вызываем функцию проверки
	err := validateUser(imya, vozrast, pochta)

	// Выводим результат
	fmt.Println("--------------------------------")
	if err != nil {
		// Если вернулась ошибка, печатаем её текст
		fmt.Println("Результат:", err)
	} else {
		// Если ошибок нет (nil)
		fmt.Println("Результат: nil (все правильно)")
	}
}
