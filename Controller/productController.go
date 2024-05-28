package controllers

import (
	models "myapp/Model"
	"myapp/ent"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := models.GetProduct(c, productID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(c *gin.Context) {
	var product ent.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, product)
		return
	}
	err = models.CreateProduct(c, &product)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": product.Name + " created successfully"})
	}
}

func UpdateProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	var product ent.Product
	err = c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, product)
		return
	}
	err = models.UpdateProduct(c, &product, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": productIDStr + " updated successfully"})
	}
}

func DeleteProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	err = models.DeleteProduct(c, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": productIDStr + " deleted successfully"})
	}
}
