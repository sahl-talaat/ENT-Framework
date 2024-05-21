package service

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
)

type ProductOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewProductOps(ctx context.Context) *ProductOps {
	return &ProductOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *ProductOps) ProductsGetAll() ([]*ent.Product, error) {

	products, err := r.client.Product.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductOps) ProductGetByID(id int) (*ent.Product, error) {

	product, err := r.client.Product.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductOps) ProductCreate(newProduct ent.Product) (*ent.Product, error) {

	newCreatedProduct, err := r.client.Product.Create().
		SetName(newProduct.Name).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedProduct, nil
}

func (r *ProductOps) ProductUpdate(product ent.Product) (*ent.Product, error) {

	updatedProduct, err := r.client.Product.UpdateOneID(product.ID).
		SetName(product.Name).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (r *ProductOps) ProductDelete(id int) (int, error) {

	err := r.client.Product.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
