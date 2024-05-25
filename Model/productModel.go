package models

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
)

func GetProduct(ctx context.Context /* product *ent.Product, */, id int) (product *ent.Product, err error) {
	product, err = config.Client.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func CreateProduct(ctx context.Context, product *ent.Product) (err error) {
	_, err = config.Client.Product.
		Create().
		SetName(product.Name).
		SetPrice(product.Price).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(ctx context.Context, product *ent.Product, id int) (err error) {
	_, err = config.Client.Product.
		UpdateOneID(id).
		SetName(product.Name).
		SetPrice(product.Price).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(ctx context.Context, id int) (err error) {
	if err = config.Client.Product.
		DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

/*
	 func ProductsGetAll() ([]*ent.Product, error) {

		products, err := config.Client.Product.Query().All(s.ctx)
		if err != nil {
			return nil, err
		}

		return products, nil
	}
*/
