package service

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
	"entdemo/ent/branch"
)

/* func (r *BranchOps) BranchCreate(newBranch ent.Branch) (*ent.Branch, error) {

} */

type BranchOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewBranchOps(ctx context.Context) *BranchOps {
	return &BranchOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *BranchOps) BranchsGetAll() ([]*ent.Branch, error) {

	users, err := r.client.Branch.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *BranchOps) BranchGetByID(id int) (*ent.Branch, error) {
	Branch, err := r.client.Branch.Query().Where(branch.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return Branch, nil
}

func (r *BranchOps) BranchCreate(newBranch ent.Branch) (*ent.Branch, error) {

	newCreatedBranch, err := r.client.Branch.Create().
		SetName(newBranch.Name).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedBranch, nil
}

func (r *BranchOps) BranchUpdate(user ent.Branch) (*ent.Branch, error) {

	updatedBranch, err := r.client.Branch.UpdateOneID(user.ID).
		SetName(user.Name).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedBranch, nil
}

func (r *BranchOps) BranchDelete(id int) (int, error) {

	err := r.client.Branch.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
