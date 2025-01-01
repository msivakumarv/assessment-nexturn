package main

import (
	"errors"
	"fmt"
	"sort"
)

type Item struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var catalog []Item

func addItem(id int, name string, price interface{}, stock int) error {
	for _, item := range catalog {
		if item.ID == id {
			return errors.New("Item ID must be unique")
		}
	}

	priceFloat, ok := price.(float64)
	if !ok {
		return errors.New("Invalid price type")
	}

	if priceFloat <= 0 {
		return errors.New("Price must be greater than zero")
	}
	if stock < 0 {
		return errors.New("Stock cannot be negative")
	}

	catalog = append(catalog, Item{
		ID:    id,
		Name:  name,
		Price: priceFloat,
		Stock: stock,
	})
	return nil
}

func updateInventory(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("Stock cannot be negative")
	}
	for i, item := range catalog {
		if item.ID == id {
			catalog[i].Stock = newStock
			return nil
		}
	}
	return errors.New("Item not found")
}

func findItemByID(id int) (Item, error) {
	for _, item := range catalog {
		if item.ID == id {
			return item, nil
		}
	}
	return Item{}, errors.New("Item not found")
}

func findItemByName(name string) (Item, error) {
	for _, item := range catalog {
		if item.Name == name {
			return item, nil
		}
	}
	return Item{}, errors.New("Item not found")
}

func showCatalog() {
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println("--------------------------------------------")
	for _, item := range catalog {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", item.ID, item.Name, item.Price, item.Stock)
	}
}

func sortByPrice() {
	sort.Slice(catalog, func(i, j int) bool {
		return catalog[i].Price < catalog[j].Price
	})
}

func sortByStock() {
	sort.Slice(catalog, func(i, j int) bool {
		return catalog[i].Stock < catalog[j].Stock
	})
}

func main() {
	addItem(1, "Laptop", 1000.00, 10)
	addItem(2, "Mouse", 25.50, 50)
	addItem(3, "Keyboard", 45.75, 30)

	fmt.Println("Initial Catalog:")
	showCatalog()

	updateInventory(2, 60)
	fmt.Println("\nCatalog after updating stock:")
	showCatalog()

	sortByPrice()
	fmt.Println("\nCatalog sorted by price:")
	showCatalog()

	sortByStock()
	fmt.Println("\nCatalog sorted by stock:")
	showCatalog()

	if item, err := findItemByID(2); err == nil {
		fmt.Println("\nItem found by ID:", item)
	} else {
		fmt.Println(err)
	}

	if item, err := findItemByName("Keyboard"); err == nil {
		fmt.Println("\nItem found by Name:", item)
	} else {
		fmt.Println(err)
	}
}
