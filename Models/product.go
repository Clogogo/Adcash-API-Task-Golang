package Models

import (
	database "AdcashTask/Database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"product_Name"`
	CatId int    `gorm:"column:catId"`
	//CategoryID int `gorm:"column:catId"`
	//Category Category `gorm:"foreignKey:CatId"`
}

func GetProducts(c *fiber.Ctx) {
	db := database.DBConn
	var product []Product
	db.Find(&product)
	_ = c.JSON(product)
}

func GetProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var product Product
	db.Find(&product, id)

	if product.Name == "" {
		c.Status(404).Send("No Product Found with ID ")
		return
	}
	_ = c.JSON(product)

}

func AddProduct(c *fiber.Ctx) {
	db := database.DBConn
	product := new(Product)

	if err := c.BodyParser(product); err != nil {
		c.Status(503).Send(err)
		return
	}

	if product.Name == "" {
		c.Status(422).Send("Invalid entry")
		return
	}

	checker := db.Where(&Product{Name: product.Name, CatId: product.CatId}).First(&product).RecordNotFound()

	if !checker {
		c.Status(409).Send("Product Already Exit with Category")
		return
	} else if checker {
		db.Create(&product)
		c.Status(200).Send("Category Successfully Added")
		_ = c.JSON(product)
	}

}

func UpdateProduct(c *fiber.Ctx) {

	db := database.DBConn

	product := new(Product)
	//db.Find(&product, id)

	if err := c.BodyParser(product); err != nil {
		db.Update(product.Name)
		return
	}

	checker := db.Where(&Product{Name: product.Name, CatId: product.CatId}).First(&product).RecordNotFound()

	if product.Name == "" {
		c.Status(404).Send("No Product found with id")
		return
	}
	if !checker {
		c.Status(409).Send("Product Already Exit with Category")
		return
	} else if checker {
		db.Save(&product)
		c.Status(200).Send("Product Updated")
	}

}

func DeleteProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var product Product
	db.Find(&product, id)

	if product.Name == "" {
		c.Status(404).Send("Product Not Found")
		return
	}

	db.Delete(&product)
	c.Send("Product Successfully Deleted")

}
