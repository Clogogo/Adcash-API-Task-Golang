package Controllers

import (
	"AdcashTask/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetProducts ... Get all categories
func GetProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}

}

//CreateProduct ... Create Product
func CreateProduct(c *gin.Context) {
	var product Models.Product

	//validate
	bindErr := c.BindJSON(&product)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//GetProductByID ... Get the Product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//UpdateProduct ... Update the user information
func UpdateProduct(c *gin.Context) {
	var product Models.Product

	bindErr := c.BindJSON(&product)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	err := Models.UpdateProduct(&product, product.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//DeleteProduct ... Delete the Product
func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
