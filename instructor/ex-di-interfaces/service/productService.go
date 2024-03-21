package service

import (
	"ex-di-interfaces/model"
	"ex-di-interfaces/repository"
)

// var repo repository.InMemoryProductRepo = repository.NewInMemoryProductRepo()

type ProductService struct {
	Repo repository.InMemoryProductRepo
	// repo repository.InMemoryProductRepo
}

func (ps ProductService) GetAllProducts() []model.Product {
	// return repo.FindAll()
	return ps.Repo.FindAll()
}

func (ps ProductService) GetProductById(id int) *model.Product {
	// repo := repository.NewInMemoryProductRepo()
	// return repo.FindBy(id)
	return ps.Repo.FindBy(id)
}

func (ps ProductService) AddProduct(name string, category model.Category, price float32) *model.Product {
	// repo := repository.NewInMemoryProductRepo()
	newProduct := model.Product{Name: name, Category: category, Price: price}

	return ps.Repo.Save(newProduct)
}

func NewProductService() ProductService {
	return ProductService{
		Repo: repository.NewInMemoryProductRepo(),
	}
}
