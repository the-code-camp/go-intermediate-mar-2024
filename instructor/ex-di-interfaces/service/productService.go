package service

import (
	"ex-di-interfaces/model"
	"ex-di-interfaces/repository"
)

// var repo repository.InMemoryProductRepo = repository.NewInMemoryProductRepo()

type ProductRepository interface {
	FindBy(id int) *model.Product
	FindAll() []model.Product
	Save(newProduct model.Product) *model.Product
}

type DefaultProductService struct {
	Repo repository.InMemoryProductRepo
	// Repo repository.DbProductRepo
}

func (ps DefaultProductService) GetAllProducts() []model.Product {
	// return repo.FindAll()
	return ps.Repo.FindAll()
}

func (ps DefaultProductService) GetProductById(id int) *model.Product {
	// repo := repository.NewInMemoryProductRepo()
	// return repo.FindBy(id)
	return ps.Repo.FindBy(id)
}

func (ps DefaultProductService) AddProduct(name string, category model.Category, price float32) *model.Product {
	// repo := repository.NewInMemoryProductRepo()
	newProduct := model.Product{Name: name, Category: category, Price: price}

	return ps.Repo.Save(newProduct)
}

func NewProductService(repo repository.InMemoryProductRepo) DefaultProductService {
	return DefaultProductService{
		Repo: repo,
	}
}
