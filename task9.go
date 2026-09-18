package main

import "fmt"

// Тип номера
type RoomType string

const (
	Single RoomType = "single"
	Double RoomType = "double"
	Suite  RoomType = "suite"
)

// Статус номера
type RoomStatus string

const (
	Free        RoomStatus = "free"
	Booked      RoomStatus = "booked"
	Maintenance RoomStatus = "maintenance"
)

type HotelRoom struct {
	Type   RoomType
	Status RoomStatus
	Price  float64
}

// Карта: номер комнаты -> структура
var nomera = map[string]HotelRoom{
	"101": {Single, Free, 2100},
	"102": {Double, Free, 3500},
	"103": {Suite, Maintenance, 6000},
}

// Функция бронирования меняет статус на booked
func zabronirovat(nomer string) {
	komnata := nomera[nomer]
	komnata.Status = Booked
	nomera[nomer] = komnata
}

func main() {
	zabronirovat("101")
	fmt.Println(nomera["101"])
}