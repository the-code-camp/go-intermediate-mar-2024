package repository

import (
	"03-lab/model"
)

type InMemoryProductRepo struct {
	products map[int]model.Product
}

func (r InMemoryProductRepo) FindBy(id int) *model.Product {
	if p, ok := r.products[id]; ok {
		return &p
	}
	return nil
}

func (r InMemoryProductRepo) FindAll() []model.Product {
	v := make([]model.Product, 0)
	for _, p := range r.products {
		v = append(v, p)
	}
	return v
}

func (r InMemoryProductRepo) Save(newProduct model.Product) *model.Product {
	id := len(r.products) + 1
	newProduct.Id = id
	r.products[id] = newProduct
	return &newProduct
}

func (r InMemoryProductRepo) Update(productToBeUpdated model.Product) bool {
	if product := r.FindBy(productToBeUpdated.Id); product != nil {
		r.products[productToBeUpdated.Id] = productToBeUpdated
		return true
	}
	return false
}

func NewInMemoryProductRepo() InMemoryProductRepo {
	inMemDb := InMemoryProductRepo{
		products: map[int]model.Product{
			1: model.NewProduct(1, "InMem - Baked goods", model.FOOD, 100.0),
			2: model.NewProduct(2, "InMem - Go programming", model.BOOKS, 55.0),
			3: model.NewProduct(3, "InMem - Soft drinks", model.BEVERAGE, 41.46),
			4: model.NewProduct(4, "InMem - Golang", model.BOOKS, 697.57),
			5: model.NewProduct(5, "InMem - Lego", model.TOYS, 295.37),
		},
	}
	return inMemDb
}
