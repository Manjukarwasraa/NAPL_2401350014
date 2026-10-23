
package main

import (
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type Summary struct {
	Count      int
	TotalValue float64
}

func addItem(items *[]Item, item Item) error {
	if item.ID <= 0 || item.Name == "" || item.Quantity < 0 || item.Price < 0 {
		return fmt.Errorf("invalid item details")
	}

	for _, v := range *items {
		if v.ID == item.ID {
			return fmt.Errorf("item ID already exists")
		}
	}

	*items = append(*items, item)
	return nil
}

func updateItem(items []Item, id, quantity int, price float64) error {
	if quantity < 0 || price < 0 {
		return fmt.Errorf("quantity and price cannot be negative")
	}

	for i := range items {
		if items[i].ID == id {
			items[i].Quantity = quantity
			items[i].Price = price
			return nil
		}
	}

	return fmt.Errorf("item not found")
}

func removeItem(items *[]Item, id int) error {
	for i, item := range *items {
		if item.ID == id {
			*items = append((*items)[:i], (*items)[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("item not found")
}

func showItems(items []Item) {
	if len(items) == 0 {
		fmt.Println("Inventory is empty")
		return
	}

	for _, item := range items {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}
}

func calculateSummary(items []Item, ch chan Summary) {
	fmt.Println("Summary goroutine started")

	s := Summary{
		Count: len(items),
	}

	for _, item := range items {
		s.TotalValue += float64(item.Quantity) * item.Price
	}

	ch <- s
	fmt.Println("Summary goroutine completed")
}

func main() {
	var items []Item

	for {
		var choice int

		fmt.Println("\n--- Inventory Management ---")
		fmt.Println("1. Add Item")
		fmt.Println("2. Update Item")
		fmt.Println("3. Remove Item")
		fmt.Println("4. Show Items")
		fmt.Println("5. Inventory Summary")
		fmt.Println("6. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var item Item

			fmt.Print("Enter ID, Name, Quantity and Price: ")
			fmt.Scan(&item.ID, &item.Name, &item.Quantity, &item.Price)

			err := addItem(&items, item)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item added successfully")
			}

		case 2:
			var id, quantity int
			var price float64

			fmt.Print("Enter item ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter new quantity and price: ")
			fmt.Scan(&quantity, &price)

			err := updateItem(items, id, quantity, price)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item updated successfully")
			}

		case 3:
			var id int
			fmt.Print("Enter item ID to remove: ")
			fmt.Scan(&id)

			err := removeItem(&items, id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item removed successfully")
			}

		case 4:
			showItems(items)

		case 5:
			ch := make(chan Summary)
			go calculateSummary(items, ch)

			s := <-ch
			fmt.Println("Total items:", s.Count)
			fmt.Printf("Total inventory value: %.2f\n", s.TotalValue)

		case 6:
			fmt.Println("Program ended")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
