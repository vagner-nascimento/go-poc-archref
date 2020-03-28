package app

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

func validateFoundSupplier(sup model.Supplier) (err error) {
	if len(sup.Id) <= 0 {
		err = errors.New("customer not found")
	}
	return err
}

func mapSupplierToUpdate(oldSup model.Supplier, newSup model.Supplier) model.Supplier {
	return model.Supplier{
		Id:             oldSup.Id,
		Name:           newSup.Name,
		DocumentNumber: newSup.DocumentNumber,
		CreditLimit:    newSup.CreditLimit,
	}
}
