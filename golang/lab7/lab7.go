package main

import "fmt"

type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
	ApplyDiscount(discount float64)
	SetPrice(newPrice float64)
	SetDescription(newDescription string)
}

type Book struct {
	Name        string
	Price       float64
	Description string
	Author      string
	ISBN        string
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) GetDescription() string {
	return b.Description
}

func (b *Book) ApplyDiscount(discount float64) {
	b.Price = b.Price - (b.Price * discount)
}

func (b *Book) SetPrice(newPrice float64) {
	b.Price = newPrice
}

func (b *Book) SetDescription(newDescription string) {
	b.Description = newDescription
}

type ElectronicDevice struct {
	Name        string
	Price       float64
	Description string
	Brand       string
	Model       string
}

func (e *ElectronicDevice) GetName() string {
	return e.Name
}

func (e *ElectronicDevice) GetPrice() float64 {
	return e.Price
}

func (e *ElectronicDevice) GetDescription() string {
	return e.Description
}

func (e *ElectronicDevice) ApplyDiscount(discount float64) {
	e.Price = e.Price - (e.Price * discount)
}

func (e *ElectronicDevice) SetPrice(newPrice float64) {
	e.Price = newPrice
}

func (e *ElectronicDevice) SetDescription(newDescription string) {
	e.Description = newDescription
}

func calculateTotalCost(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.GetPrice()
	}
	return totalCost
}

func main() {
	book1 := &Book{
		Name:        "The Hitchhiker's Guide to the Galaxy",
		Price:       12.99,
		Description: "A humorous science fiction novel",
		Author:      "Douglas Adams",
		ISBN:        "0-345-39180-2",
	}

	phone := &ElectronicDevice{
		Name:        "iPhone 14 Pro Max",
		Price:       1099.00,
		Description: "A premium smartphone",
		Brand:       "Apple",
		Model:       "iPhone 14 Pro Max",
	}

	products := []Product{book1, phone}

	fmt.Printf("Общая стоимость товаров: %.2f\n", calculateTotalCost(products))

	for _, product := range products {
		if _, ok := product.(*Book); ok {
			product.ApplyDiscount(0.10)
		}
	}

	fmt.Printf("Общая стоимость товаров после скидки: %.2f\n", calculateTotalCost(products))

	phone.SetPrice(999.00)
	fmt.Printf("Новая цена телефона: %.2f\n", phone.GetPrice())

	book1.SetDescription("A humorous science fiction novel by Douglas Adams")
	fmt.Printf("Новое описание книги: %s\n", book1.GetDescription())
}
