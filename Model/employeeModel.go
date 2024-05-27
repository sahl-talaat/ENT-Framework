package models

import (
	"context"
	config "entdemo/Config"
	"entdemo/ent"
	"entdemo/ent/employee"
	"time"
)

func GetEmployees(ctx context.Context) (employees []*ent.Employee, err error) {
	employees, err = config.Client.Employee.Query().WithDepartment().WithBranch().All(ctx)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func CreateEmployee(ctx context.Context, name string, age int, salary float64, departmentID, branchID int) (err error) {
	currentTime := time.Now()
	_, err = config.Client.Employee.
		Create().
		SetName(name).
		SetAge(age).
		SetSalary(salary).
		SetCreatedAt(currentTime).
		SetDepartmentID(departmentID).
		SetBranchID(branchID).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetEmployee(ctx context.Context, id int) (emp *ent.Employee, err error) {
	emp, err = config.Client.Employee.Query().
		WithDepartment().
		WithBranch().
		Where(employee.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return emp, nil
}

func UpdateEmployee(ctx context.Context, id int, name string, age int, salary float64, departmentID, branchID int) (err error) {
	_, err = config.Client.Employee.
		UpdateOneID(id).
		SetName(name).
		SetAge(age).
		SetSalary(salary).
		SetDepartmentID(departmentID).
		SetBranchID(branchID).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEmployee(ctx context.Context, id int) (err error) {
	if err = config.Client.Employee.DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil
}
