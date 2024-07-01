package main

import (
	"fmt"
	"strings"
)

type LibraryItem struct {
	Title string
	Type  string
}

func (li LibraryItem) Info() string {
	return fmt.Sprintf("Title: %s, Type: %s", li.Title, li.Type)
}

type Library struct {
	Items []LibraryItem
}

func (lib *Library) AddItem(item LibraryItem) {
	lib.Items = append(lib.Items, item)
}

func (lib Library) SearchItem(title string) string {
	for _, item := range lib.Items {
		if strings.EqualFold(item.Title, title) {
			return item.Info()
		}
	}
	return "Item not found"
}

func (lib Library) ListItems() {
	for _, item := range lib.Items {
		fmt.Println(item.Info())
	}
}

// Define an interface for searchable items
type Searchable interface {
	SearchItem(title string) string
}

// Function to search any searchable collection
func searchCollection(s Searchable, title string) {
	fmt.Println(s.SearchItem(title))
}

func main() {
	library := Library{}

	// Create some library items
	book1 := LibraryItem{Title: "Go Programming in Go", Type: "Book"}
	book2 := LibraryItem{Title: "Learning Go", Type: "Book"}
	magazine2 := LibraryItem{Title: "Science Today", Type: "Magazine"}
	magazine1 := LibraryItem{Title: "Science Tomorrow", Type: "Magazine"}

	// Add items to the library
	library.AddItem(book1)
	library.AddItem(book2)
	library.AddItem(magazine1)
	library.AddItem(magazine2)

	fmt.Println("Library Items:")
	library.ListItems()

	// Search for items in the library
	fmt.Println("\nSearch Results:")
	searchCollection(library, "Learning Go")
	searchCollection(library, "Science Today")
	searchCollection(library, "All the Science")
}
