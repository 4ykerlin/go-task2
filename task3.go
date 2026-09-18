package main

import (
	"fmt"
)


type Order struct {
	ID          int    
	Items       []int  
	Total       int    
	Address     string 
	IsCompleted bool   
}


func AddOrder(orders map[int]Order, newOrder Order) {
	orders[newOrder.ID] = newOrder
}

func main() {

	ordersMap := make(map[int]Order)

	
	order1 := Order{
		ID:          1001,
		Items:       []int{105, 204, 312},
		Total:       150050, 
		Address:     "г. Москва, ул. Ленина, д. 1, кв. 10",
		IsCompleted: false,
	}

	
	order2 := Order{
		ID:          1002,
		Items:       []int{405, 112},
		Total:       89000, 
		Address:     "г. Санкт-Петербург, Невский пр., д. 50",
		IsCompleted: true,
	}

	
	AddOrder(ordersMap, order1)
	AddOrder(ordersMap, order2)


	fmt.Println("--- Список заказов в магазине ---")
	

	for id, order := range ordersMap {
		rubles := order.Total / 100
		kopecks := order.Total % 100
		
		fmt.Printf("ID заказа: %d\n", id)
		fmt.Printf("  Товары: %v\n", order.Items)
		fmt.Printf("  Сумма: %d руб. %02d коп.\n", rubles, kopecks)
		fmt.Printf("  Адрес: %s\n", order.Address)
		fmt.Printf("  Статус: %t\n\n", order.IsCompleted)
	}
}