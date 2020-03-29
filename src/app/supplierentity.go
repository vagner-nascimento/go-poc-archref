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
		IsActive:       newSup.IsActive,
	}
}

func validateEnterprise(ent model.Enterprise) (err error) {
	if len(ent.Document) <= 0 {
		err = errors.New("model.Enterprise document is required")
	}
	return err
}

func mapSupplierFromEnterprise(foundSup model.Supplier, ent model.Enterprise) model.Supplier {
	return model.Supplier{
		Id:             foundSup.Id,
		Name:           ent.CompanyName,
		DocumentNumber: ent.Document,
		IsActive:       ent.Active,
		CreditLimit:    foundSup.CreditLimit,
	}
}
