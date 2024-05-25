package controller

import (
	models "entdemo/Model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEmployee(c *gin.Context) {
	employeeIDStr := c.Param("id")
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}
	employee, err := models.GetEmployee(c, employeeID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func CreateEmployee(c *gin.Context) {
	//var employee ent.Employee
	var employee struct {
		Name         string  `json:"name"`
		Age          int     `json:"age"`
		Salary       float64 `json:"salary"`
		DepartmentID int     `json:"department_id"`
		BranchID     int     `json:"branch_id"`
	}
	err := c.BindJSON(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, employee)
		return
	}
	err = models.CreateEmployee(c, employee.Name, employee.Age, employee.Salary, employee.DepartmentID, employee.BranchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": employee.Name + " created successfully"})
	}
}

func UpdateEmployee(c *gin.Context) {
	employeeIDStr := c.Param("id")
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}
	var employee struct {
		Name         string  `json:"name"`
		Age          int     `json:"age"`
		Salary       float64 `json:"salary"`
		DepartmentID int     `json:"department_id"`
		BranchID     int     `json:"branch_id"`
	}
	err = c.BindJSON(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, employee)
		return
	}
	err = models.UpdateEmployee(c, employeeID, employee.Name, employee.Age, employee.Salary, employee.DepartmentID, employee.BranchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": employeeIDStr + " updated successfully"})
	}
}

func DeleteEmployee(c *gin.Context) {
	employeeIDStr := c.Param("id")
	employeeID, err := strconv.Atoi(employeeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}
	err = models.DeleteEmployee(c, employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": employeeIDStr + " deleted successfully"})
	}
}
