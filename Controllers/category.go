package Controllers

import (
	"AdcashTask/Models"
	"AdcashTask/Utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetCategories ... Get all categories
func GetCategories(c *gin.Context) {
	var categories []Models.Category
	err := Models.GetAllCategories(&categories)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, categories)
	}
}

//CreateCategory ... Create category
func CreateCategory(c *gin.Context) {
	var category Models.Category
	bindErr := c.BindJSON(&category)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}
	err := Models.CreateCategory(&category)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusConflict, "Duplicate Entry")
		return
	} else {
		c.JSON(http.StatusOK, category)
	}
}

//Get category Product By ID
func GetCategoryProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := Models.CategoryProductList(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Category Not Found")
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//GetCategoryByID ... Get the Category by id
func GetCategoryByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var category Models.Category

	err := Models.GetCategoryByID(&category, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Category Not Found")
	} else {
		c.JSON(http.StatusOK, category)
	}
}

//UpdateCategory ... Update the user information
func UpdateCategory(c *gin.Context) {
	var category Models.Category

	// App level validation
	bindErr := c.BindJSON(&category)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Check if resource exist
	err := Models.UpdateCategory(&category, category.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}

}

//DeleteCategory ... Delete the category
func DeleteCategory(c *gin.Context) {
	var category Models.Category
	id := Utilities.GetInt64IdFromReqContext(c)

	err := Models.DeleteCategory(&category, id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Category Not Found")
	} else {
		c.JSON(http.StatusOK, "Deleted")
	}

}
