package models

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
)

func GetDepartments(ctx context.Context) (departments []*ent.Department, err error) {
	departments, err = config.Client.Department.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func CreateDepartment(ctx context.Context, department *ent.Department) (err error) {
	_, err = config.Client.Department.
		Create().
		SetName(department.Name).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetDepartment(ctx context.Context, id int) (department *ent.Department, err error) {
	department, err = config.Client.Department.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func UpdateDepartment(ctx context.Context, department *ent.Department, id int) (err error) {
	_, err = config.Client.Department.
		UpdateOneID(id).
		SetName(department.Name).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDepartment(ctx context.Context, id int) (err error) {
	if err = config.Client.Department.
		DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

/* func DepartmentsGetAll() ([]*ent.Department, error) {

	users, err := r.client.Department.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
} */
