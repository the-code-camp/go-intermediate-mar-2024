package service

import "ex-di-interfaces/model"

/**
Abstraction for repository can also be structured like this. This makes more sense and aligns with
Go's philosophy. Important thing is smaller focussed interfaces which can be used composed for larger
functionality
*/

type DataFinder interface {
	FindBy(id int) *model.Product
	FindAll() []model.Product
}

type DataSaver interface {
	Save(newProduct model.Product) *model.Product
}

type DataUpdater interface {
	Update(productToBeUpdated model.Product) bool
}

// Below two are examples of composing interfaces for more functionality
type DataFindSave interface {
	DataFinder
	DataSaver
}

type DataFindSaveUpdate interface {
	DataFinder
	DataSaver
	DataUpdater
}
