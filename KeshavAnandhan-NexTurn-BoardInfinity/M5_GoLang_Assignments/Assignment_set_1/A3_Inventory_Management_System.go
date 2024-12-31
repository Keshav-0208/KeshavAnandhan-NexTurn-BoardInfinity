package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type product struct {
	id    int
	name  string
	price float64
	stock int
}

var inventory = []product{
	{1, "PlayStation 5", 499.99, 15},
	{2, "Xbox Series X", 499.99, 10},
	{3, "Nintendo Switch", 299.99, 20},
}

func add_product(id int, name string, price float64, stock int) error {
	for _, p := range inventory {
		if p.id == id {
			return errors.New("this product ID already exists")
		}
	}
	if price <= 0 {
		return errors.New("price must be positive")
	}
	if stock < 0 {
		return errors.New("stock can't be negative")
	}
	inventory = append(inventory, product{id, name, price, stock})
	return nil
}

func update_stock(id, stock int) error {
	if stock < 0 {
		return errors.New("stock can't be negative")
	}
	for i, p := range inventory {
		if p.id == id {
			inventory[i].stock = stock
			return nil
		}
	}
	return errors.New("product not found")
}

func search_product(term string) (*product, error) {
	for _, p := range inventory {
		if strings.EqualFold(p.name, term) || fmt.Sprintf("%d", p.id) == term {
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}

func display_inventory() {
	fmt.Printf("\n%-5s %-20s %-10s %-10s\n", "ID", "Product", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, p := range inventory {
		fmt.Printf("%-5d %-20s $%-9.2f %-10d\n", p.id, p.name, p.price, p.stock)
	}
}

func sort_inventory(by string) {
	switch by {
	case "price":
		sort.Slice(inventory, func(i, j int) bool { return inventory[i].price < inventory[j].price })
	case "stock":
		sort.Slice(inventory, func(i, j int) bool { return inventory[i].stock < inventory[j].stock })
	default:
		fmt.Println("Sort by 'price' or 'stock' only.")
	}
}

func main() {
	for {
		fmt.Println("\nNexTurn Electronics")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Show Inventory")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		var choice, id, stock int
		var name, term, sort_by string
		var price float64
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting... Have a great day!")
			break
		}

		switch choice {
		case 1:
			fmt.Print("Enter ID, Name, Price, Stock: ")
			fmt.Scan(&id, &name, &price, &stock)
			err := add_product(id, name, price, stock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(name, "has been added to the inventory.")
			}

		case 2:
			fmt.Print("Enter Product ID and New Stock: ")
			fmt.Scan(&id, &stock)
			err := update_stock(id, stock)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully!")
			}

		case 3:
			fmt.Print("Enter Product Name or ID: ")
			fmt.Scan(&term)
			p, err := search_product(term)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product Found: ID=%d, Name=%s, Price=$%.2f, Stock=%d\n", p.id, p.name, p.price, p.stock)
			}

		case 4:
			display_inventory()

		case 5:
			fmt.Print("Sort by 'price' or 'stock': ")
			fmt.Scan(&sort_by)
			sort_inventory(sort_by)
			display_inventory()

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
