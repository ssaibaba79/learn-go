package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	Quantity int    `json:"quantity"`
}

type ItemList struct {
	Name  string `json:"Name"`
	Items []Item `json:"items"`
}

func (list *ItemList) addItem(item Item) {
	list.Items = append(list.Items, item)
}

func (list ItemList) getItem(id int) (Item, error) {
	for _, v := range list.Items {
		if v.Id == id {
			return v, nil
		}
	}

	return Item{}, fmt.Errorf("Item with Id %d not found", id)
}

func (list ItemList) deleteItem(id int) (err error) {
	var index int = -1
	for i, v := range list.Items {
		if v.Id == id {
			index = i
		}
	}
	if index >= 0 {
		list.Items = append(list.Items[:index], list.Items[index+1:]...)
		return
	} else {
		return fmt.Errorf("Item with Id %d not found", id)
	}
}

func toJson(v any) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func main() {

	var item1 = Item{Id: 1000}
	item1.Name = "Banana"
	item1.Quantity = 5

	var item2 = Item{1001, "Milk", "Gallon(s)", 1}
	var list = ItemList{Name: "To buy", Items: []Item{}}
	fmt.Println(toJson(list))

	var item3 = Item{1002, "Rice", "lbs)", 10}

	list.addItem(item1)
	list.addItem(item2)
	list.addItem(item3)
	fmt.Println(toJson(list))

	item, err := list.getItem(item1.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(toJson(item))
	}

	item, err = list.getItem(5000)
	if err != nil {
		fmt.Println(err)
	}

	list.deleteItem(item2.Id)
	fmt.Println(toJson(list))

}
