package main

// go mod init github.com/chrismarsilva/cms.project.terminal.widget
// go get -u github.com/rivo/tview@master
// go mod tidy

// go run main.go

// go install github.com/air-verse/air@latest
// air init
// air

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/rivo/tview"
)

var (
	inventory     = []Item{}         // Our inventory list, initially empty
	inventoryFile = "inventory.json" // File where inventory will be saved/loaded from
)

func main() {
	app := tview.NewApplication()

	loadInventory()

	inventoryList := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	inventoryList.SetBorder(true).SetTitle("Inventory Items")

	refreshInventory := func() {
		inventoryList.Clear()

		if len(inventory) == 0 {
			fmt.Fprintln(inventoryList, "No items in inventory.")
		} else {
			for i, item := range inventory {
				fmt.Fprintf(inventoryList, "[%d] %s (Stock: %d)\n", i+1, item.Name, item.Stock)
			}
		}
	}

	itemNameInput := tview.NewInputField().SetLabel("Item Name: ")
	itemStockInput := tview.NewInputField().SetLabel("Stock: ")
	itemIDInput := tview.NewInputField().SetLabel("Item ID to delete: ")

	form := tview.NewForm().
		AddFormItem(itemNameInput).
		AddFormItem(itemStockInput).
		AddFormItem(itemIDInput).
		AddButton("Add Item", func() {
			name := itemNameInput.GetText()
			stock := itemStockInput.GetText()

			if name != "" && stock != "" {
				quantity, err := strconv.Atoi(stock)
				if err != nil {
					fmt.Fprintln(inventoryList, "Invalid stock value.")
					return
				}

				inventory = append(inventory, Item{Name: name, Stock: quantity})
				saveInventory()
				refreshInventory()
				itemNameInput.SetText("")
				itemStockInput.SetText("")
			}
		}).
		AddButton("Delete Item", func() {
			idStr := itemIDInput.GetText()
			if idStr == "" {
				fmt.Fprintln(inventoryList, "Please enter an item ID to delete.")
				return
			}

			id, err := strconv.Atoi(idStr)
			if err != nil || id < 1 || id > len(inventory) {
				fmt.Fprintln(inventoryList, "Invalid item ID.")
				return
			}

			deleteItem(id - 1)
			fmt.Fprintf(inventoryList, "Item [%d] deleted.\n", id)
			refreshInventory()
			itemIDInput.SetText("")
		}).
		AddButton("Exit", func() { // Button to exit the application
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Manage Inventory").SetTitleAlign(tview.AlignLeft)

	flex := tview.NewFlex().
		AddItem(inventoryList, 0, 1, false).
		AddItem(form, 0, 1, true)

	refreshInventory()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		slog.Error("Error running program", slog.Any("err", err))
	}

	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")

	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// 	slog.Error("Error running program", slog.Any("err", err))
	// }
}

type Item struct {
	Name  string `json:"name"`  // Name of the item (will be stored as JSON)
	Stock int    `json:"stock"` // Quantity of the item in stock (also stored as JSON)
}

func loadInventory() {
	if _, err := os.Stat(inventoryFile); err == nil {
		data, err := os.ReadFile(inventoryFile)
		if err != nil {
			log.Fatal("Error reading inventory file:", err)
		}

		json.Unmarshal(data, &inventory)
	}
}

func saveInventory() {
	data, err := json.MarshalIndent(inventory, "", "  ")
	if err != nil {
		log.Fatal("Error saving inventory:", err)
	}

	os.WriteFile(inventoryFile, data, 0644)
}

func deleteItem(index int) {
	if index < 0 || index >= len(inventory) {
		fmt.Println("Invalid item index.")
		return
	}

	inventory = append(inventory[:index], inventory[index+1:]...)
	saveInventory()
}
