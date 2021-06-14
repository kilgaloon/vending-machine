package model

import (
	"errors"

	"github.com/kilgaloon/atm/utils"
	"gorm.io/gorm"
)

var (
	ErrProductAlreadyExist = errors.New("Product with same name already exist")
	ErrProductNotFound = errors.New("Product not found")
)

type Product struct {
	gorm.Model
	AmountAvailable uint64 `json:"amount_available"`
	Cost            uint64 `json:"cost"`
	ProductName     string `json:"product_name"`
	SellerID        uint   `json:"seller_id"`
}

func (p *Product) Find() (*Product, error) {
	db := utils.DBConnect()

	// find product
	result := db.Where("(product_name = ? AND seller_id = ?) OR id = ?", p.ProductName, p.SellerID, p.Model.ID).Find(&Product{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrProductNotFound
	}

	return p, nil
}

func (p *Product) Create() (*Product, error) {
	db := utils.DBConnect()

	result := db.Where("product_name = ? AND seller_id = ?", p.ProductName, p.SellerID).Find(&Product{})
	if result.RowsAffected > 0 {
		return nil, ErrProductAlreadyExist
	}

	result = db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}


func (p *Product) Update() (*Product, error) {
	db := utils.DBConnect()

	result := db.Save(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}


func (p *Product) Delete() error {
	db := utils.DBConnect()

	// find user by provided username
	result, err := p.Find()
	if err != nil {
		return err
	}

	deleteResult := db.Delete(result)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
