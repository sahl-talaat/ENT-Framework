package service

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
	"entdemo/ent/employee"
)

/* func (r *EmployeeOps) EmployeeCreate(newEmployee ent.Employee) (*ent.Employee, error) {

} */

type EmployeeOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewEmployeeOps(ctx context.Context) *EmployeeOps {
	return &EmployeeOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *EmployeeOps) EmployeesGetAll() ([]*ent.Employee, error) {

	users, err := r.client.Employee.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *EmployeeOps) EmployeeGetByID(id int) (*ent.Employee, error) {
	employee, err := r.client.Employee.Query().Where(employee.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *EmployeeOps) EmployeeCreate(newEmployee ent.Employee) (*ent.Employee, error) {

	newCreatedEmployee, err := r.client.Employee.Create().
		SetEmail(newEmployee.Email).
		SetName(newEmployee.Name).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedEmployee, nil
}

func (r *EmployeeOps) EmployeeUpdate(user ent.Employee) (*ent.Employee, error) {

	updatedEmployee, err := r.client.Employee.UpdateOneID(user.ID).
		SetEmail(user.Email).
		SetName(user.Name).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedEmployee, nil
}

func (r *EmployeeOps) EmployeeDelete(id int) (int, error) {

	err := r.client.Employee.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
