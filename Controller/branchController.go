package controller

import (
	models "entdemo/Model"
	"entdemo/ent"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBranch(c *gin.Context) {
	branchIDStr := c.Param("id")
	branchID, err := strconv.Atoi(branchIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
		return
	}
	branch, err := models.GetBranch(c, branchID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, branch)
	}
}

func CreateBranch(c *gin.Context) {
	var branch ent.Branch
	err := c.BindJSON(&branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, branch)
		return
	}
	err = models.CreateBranch(c, &branch)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": branch.Name + " created successfully"})
	}
}

func UpdateBranch(c *gin.Context) {
	branchIDStr := c.Param("id")
	branchID, err := strconv.Atoi(branchIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
		return
	}
	var branch ent.Branch
	err = c.BindJSON(&branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, branch)
		return
	}
	err = models.UpdateBranch(c, &branch, branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": branchIDStr + " updated successfully"})
	}
}

func DeleteBranch(c *gin.Context) {
	branchIDStr := c.Param("id")
	branchID, err := strconv.Atoi(branchIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
		return
	}
	err = models.DeleteBranch(c, branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": branchIDStr + " deleted successfully"})
	}
}
