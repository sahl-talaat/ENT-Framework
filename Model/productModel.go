package models

import (
	"context"
	config "myapp/Config"
	"myapp/ent"
)

func GetProducts(ctx context.Context) (products []*ent.Product, err error) {
	products, err = config.Client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
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

func GetProduct(ctx context.Context /* product *ent.Product, */, id int) (product *ent.Product, err error) {
	product, err = config.Client.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
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
