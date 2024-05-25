package models

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
)

func GetBranch(ctx context.Context, id int) (branch *ent.Branch, err error) {
	branch, err = config.Client.Branch.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return branch, nil
}

func CreateBranch(ctx context.Context, branch *ent.Branch) (err error) {
	_, err = config.Client.Branch.Create().SetName(branch.Name).SetAddress(branch.Address).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBranch(ctx context.Context, branch *ent.Branch, id int) (err error) {
	_, err = config.Client.Branch.
		UpdateOneID(id).
		SetName(branch.Name).
		SetAddress(branch.Address).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBranch(ctx context.Context, id int) (err error) {
	if err = config.Client.Branch.
		DeleteOneID(id).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

/*
func BranchsGetAll(branches []*ent.Branch) (err error) {

	users, err := r.client.Branch.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}*/
