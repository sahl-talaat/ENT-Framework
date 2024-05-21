package service

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
)

type OrderOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewOrderOps(ctx context.Context) *OrderOps {
	return &OrderOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}
