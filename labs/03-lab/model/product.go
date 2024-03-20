package model

import "fmt"

type Product struct {
	Id       int
	Name     string
	Category Category
	Price    float32
}

func (p Product) String() string {
	return fmt.Sprintf("{ id:%d, name:'%s', category:'%s', price:%5.2f }", p.Id, p.Name, p.Category, p.Price)
}

func NewProduct(id int, name string, category Category, price float32) Product {
	return Product{id, name, category, price}
}
