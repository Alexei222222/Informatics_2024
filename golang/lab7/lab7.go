package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
	ApplyDiscount(discount float64)
	SetPrice(newPrice float64)
	SetDescription(newDescription string)
	GetOldPrice() float64
}

type Book struct {
	Name        string
	Price       float64
	OldPrice    float64
	Description string
	Author      string
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
	b.OldPrice = b.Price
	b.Price = b.Price - (b.Price * discount)
}

func (b *Book) SetPrice(newPrice float64) {
	b.OldPrice = b.Price
	b.Price = newPrice
}

func (b *Book) SetDescription(newDescription string) {
	b.Description = newDescription
}

func (b *Book) GetOldPrice() float64 {
	return b.OldPrice
}

type ElectronicDevice struct {
	Name        string
	Price       float64
	OldPrice    float64
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
	e.OldPrice = e.Price
	e.Price = e.Price - (e.Price * discount)
}

func (e *ElectronicDevice) SetPrice(newPrice float64) {
	e.OldPrice = e.Price
	e.Price = newPrice
}

func (e *ElectronicDevice) SetDescription(newDescription string) {
	e.Description = newDescription
}

func (e *ElectronicDevice) GetOldPrice() float64 {
	return e.OldPrice
}

func calculateTotalCost(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.GetPrice()
	}
	return totalCost
}

func RunLab7() {
	book1 := &Book{
		Name:        "Война и мир",
		Price:       1300,
		Description: "роман",
		Author:      "Лев Николаевич Толстой",
	}

	phone := &ElectronicDevice{
		Name:        "iPhone 16",
		Price:       140000,
		Description: "новый смартфон",
		Brand:       "Apple",
		Model:       "iPhone 16",
	}

	products := []Product{book1, phone}

	fmt.Printf("Общая стоимость товаров: %.2f\n", calculateTotalCost(products))

	for _, product := range products {
		if _, ok := product.(*Book); ok {
			product.ApplyDiscount(0.10)
			fmt.Printf("Книга '%s': Старая цена - %.2f, Новая цена - %.2f\n", product.GetName(), product.GetOldPrice(), product.GetPrice())
		}
		if _, ok := product.(*ElectronicDevice); ok {
			product.ApplyDiscount(0.05)
			fmt.Printf("Телефон '%s': Старая цена - %.2f, Новая цена - %.2f\n", product.GetName(), product.GetOldPrice(), product.GetPrice())
		}
	}

	fmt.Printf("Общая стоимость товаров после скидки: %.2f\n", calculateTotalCost(products))

	phone.SetPrice(125000)
	fmt.Printf("Телефон '%s': Старая цена - %.2f, Новая цена - %.2f\n", phone.GetName(), phone.GetOldPrice(), phone.GetPrice())

	book1.SetDescription("саамый популярный роман")
	fmt.Printf("Новое описание книги: %s\n", book1.GetDescription())
}
