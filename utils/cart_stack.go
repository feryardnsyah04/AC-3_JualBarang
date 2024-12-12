package utils

import (
	"shopping-cart-api/database"
	"shopping-cart-api/models"
	"fmt"
)

type Stack struct {
	items []models.CartItem
}

func (s *Stack) Push(item models.CartItem) (int64, error) {
	s.items = append(s.items, item)
	query := "INSERT INTO cart_items (product, variant, price, quantity) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, item.Product, item.Variant, item.Price, item.Quantity)
	if err != nil {
		fmt.Println("Error inserting item to database:", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		return 0, err
	}
	item.ID = int(id)
	return id, nil
}

func (s *Stack) Pop() (models.CartItem, bool) {
	if len(s.items) == 0 {
		return models.CartItem{}, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	_, err := database.DB.Exec("DELETE FROM cart_items WHERE id = ?", item.ID)
	if err != nil {
		fmt.Println("Error deleting item from database:", err)
	}
	return item, true
}

func (s *Stack) Top() (models.CartItem, bool) {
	if len(s.items) == 0 {
		return models.CartItem{}, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) GetAll() []models.CartItem {
	return s.items
}

func (s *Stack) Update(id int, variant string, quantity int) bool {
	for i, item := range s.items {
		if item.ID == id {
			s.items[i].Variant = variant
			s.items[i].Quantity = quantity
			return true
		}
	}
	return false
}

func (s *Stack) RemoveByID(id int) bool {
	index := -1
	for i, item := range s.items {
		if item.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return false
	}
	s.items = append(s.items[:index], s.items[index+1:]...)
	return true
}
