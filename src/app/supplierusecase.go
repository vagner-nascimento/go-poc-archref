package app

import (
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

type SupplierUseCase interface {
	Create(supplier *model.Supplier) error
	Find(id string) (model.Supplier, error)
	Update(id string, supplier model.Supplier) (model.Supplier, error)
	List(params []model.SearchParameter, page int64, quantity int64) ([]model.Supplier, int64, error)
}

type supplierUseCase struct {
	repository SupplierDataHandler
}

func (su *supplierUseCase) getValidSupplier(id string) (sup model.Supplier, err error) {
	if sup, err = su.repository.Get(id); err == nil {
		err = validateFoundSupplier(sup)
	}
	return sup, err
}

func (su *supplierUseCase) Create(sup *model.Supplier) error {
	return su.repository.Save(sup)
}

func (su *supplierUseCase) Find(id string) (model.Supplier, error) {
	return su.getValidSupplier(id)
}

func (su *supplierUseCase) Update(id string, sup model.Supplier) (updatedSup model.Supplier, err error) {
	if foundSup, err := su.getValidSupplier(id); err == nil {
		updatedSup = mapSupplierToUpdate(foundSup, sup)
		if err = su.repository.Update(updatedSup); err != nil {
			updatedSup = model.Supplier{}
		}
	}
	return updatedSup, err
}

func (su *supplierUseCase) List(params []model.SearchParameter, page int64, quantity int64) ([]model.Supplier, int64, error) {
	return su.repository.GetMany(params, page, quantity)
}

func NewSupplierUseCase(repository SupplierDataHandler) SupplierUseCase {
	return &supplierUseCase{repository: repository}
}
