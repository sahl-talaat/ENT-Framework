package controller

import (
	models "entdemo/Model"
	"entdemo/ent"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDepartments(c *gin.Context) {
	departments, err := models.GetDepartments(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error)
		return
	} else {
		c.JSON(http.StatusOK, departments)
	}
}

func CreateDepartment(c *gin.Context) {
	var department ent.Department
	err := c.BindJSON(&department)
	if err != nil {
		c.JSON(http.StatusBadRequest, department)
		return
	}
	err = models.CreateDepartment(c, &department)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": department.Name + " created successfully"})
	}
}

func GetDepartment(c *gin.Context) {
	departmentIDStr := c.Param("id")
	departmentID, err := strconv.Atoi(departmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	department, err := models.GetDepartment(c, departmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, department)
	}
}

func UpdateDepartment(c *gin.Context) {
	departmentIDStr := c.Param("id")
	departmentID, err := strconv.Atoi(departmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	var department ent.Department
	err = c.BindJSON(&department)
	if err != nil {
		c.JSON(http.StatusBadRequest, department)
		return
	}
	err = models.UpdateDepartment(c, &department, departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": departmentIDStr + " updated successfully"})
	}
}

func DeleteDepartment(c *gin.Context) {
	departmentIDStr := c.Param("id")
	departmentID, err := strconv.Atoi(departmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	err = models.DeleteDepartment(c, departmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": departmentIDStr + " deleted successfully"})
	}
}
