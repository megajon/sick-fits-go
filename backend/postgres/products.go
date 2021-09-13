package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/megajon/sick-fits-go/graph/model"
)

type ProductsRepo struct {
	DB *pg.DB
}

func (t *ProductsRepo) GetProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := t.DB.Model(&products).Select()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (t *ProductsRepo) CreateProduct(product *model.Product) (*model.Product, error) {
	_, err := t.DB.Model(product).Returning("*").Insert()
	return product, err
}
