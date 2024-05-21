package service

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
	"entdemo/ent/department"
)

/* func (r *DepartmentOps) DepartmentCreate(newDepartment ent.Department) (*ent.Department, error) {

} */

type DepartmentOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewDepartmentOps(ctx context.Context) *DepartmentOps {
	return &DepartmentOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *DepartmentOps) DepartmentsGetAll() ([]*ent.Department, error) {

	users, err := r.client.Department.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *DepartmentOps) DepartmentGetByID(id int) (*ent.Department, error) {
	Department, err := r.client.Department.Query().Where(department.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return Department, nil
}

func (r *DepartmentOps) DepartmentCreate(newDepartment ent.Department) (*ent.Department, error) {

	newCreatedDepartment, err := r.client.Department.Create().
		SetName(newDepartment.Name).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedDepartment, nil
}

func (r *DepartmentOps) DepartmentUpdate(user ent.Department) (*ent.Department, error) {

	updatedDepartment, err := r.client.Department.UpdateOneID(user.ID).
		SetName(user.Name).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedDepartment, nil
}

func (r *DepartmentOps) DepartmentDelete(id int) (int, error) {

	err := r.client.Department.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
