package main

import "testing"

func TestAddItem(t *testing.T) {
	var items []Item

	err := addItem(&items, Item{1, "Laptop", 2, 50000})
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if len(items) != 1 {
		t.Error("Expected 1 item")
	}
}

func TestUpdateItem(t *testing.T) {
	items := []Item{{1, "Laptop", 2, 50000}}

	err := updateItem(items, 1, 5, 45000)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if items[0].Quantity != 5 || items[0].Price != 45000 {
		t.Error("Item was not updated correctly")
	}
}

func TestRemoveItem(t *testing.T) {
	items := []Item{{1, "Laptop", 2, 50000}}

	err := removeItem(&items, 1)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	if len(items) != 0 {
		t.Error("Item was not removed")
	}
}