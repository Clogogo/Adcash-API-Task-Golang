package Routes
import (

	"AdcashTask/Models"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	app := r.Group("/api/v1")



	app.GET("Category", Models.GetCategories)
	app.GET("Category/:id", Models.GetCategory)
	app.POST("Category", Models.AddCategory)
	app.DELETE("Category/:id", Models.DeleteCategories)
	app.PUT("Category/:id", Models.UpdateCategories)

	//product routes
	app.GET("Products", Models.GetProducts)
	app.GET("Product/:id", Models.GetProduct)
	app.POST("Product", Models.AddProduct)
	app.DELETE("Product/:id", Models.DeleteProduct)
	app.PUT("Product/:id", Models.UpdateProduct)
	return r
}