package Routes

import (
	"AdcashTask/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	app := r.Group("/api/v1")

	app.GET("category", Controllers.GetCategories)
	app.GET("category/:id", Controllers.GetCategoryByID)
	app.GET("category/:id/products", Controllers.GetCategoryProductByID)
	app.POST("category", Controllers.CreateCategory)
	app.DELETE("category/:id", Controllers.DeleteCategory)
	app.PUT("category", Controllers.UpdateCategory)

	//product routes
	app.GET("products", Controllers.GetProducts)
	app.GET("product/:id", Controllers.GetProductByID)
	app.POST("product", Controllers.CreateProduct)
	app.DELETE("product/:id", Controllers.DeleteProduct)
	app.PUT("product", Controllers.UpdateProduct)
	return r
}
