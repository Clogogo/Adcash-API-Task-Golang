package Models

import (
	database "AdcashTask/Database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `validate:"required,string" json:"catName"`
	Products []Product `gorm:"ForeignKey:CatId"`
}

func (category *Category) TableName() string {
	return "category"
}


func GetCategories(c *fiber.Ctx) {
	db := database.DBConn
	var category []Category
	db.Find(&category)
	_ = c.JSON(category)
}

func GetCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var category Category
	db.Find(&category, id)

	if category.Name == "" {
		c.Status(404).Send("Category Not Found")
		return
	} else {

		_ = c.JSON(category)

	}
}

func AddCategory(c *fiber.Ctx) {
	db := database.DBConn

	category := new(Category)

	if err := c.BodyParser(category); err != nil {
		c.Status(503).Send(err)
		return
	}

	if category.Name == "" {
		c.Status(422).Send("Invalid entry")
		return
	}

	checker := db.Find(&category, "name = ?", category.Name).RecordNotFound()

	if !checker {
		c.Status(409).Send("Category Already Exit")
		return
	} else if checker {
		db.Create(&category)
		_ = c.JSON(category)
		return
	} else {
		c.Status(400)
	}

}

func DeleteCategories(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var category Category
	db.Find(&category, id)

	if category.Name == "" {
		c.Status(404).Send("No Category Found with given ID")
		return
	}

	db.Delete(&category)
	c.Send("Category Successfully Deleted")

}

func UpdateCategories(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	category := new(Category)
	db.Find(&category, id)
	if category.Name == "" {
		c.Status(404).Send("No Category found with id")
		return
	}
	if err := c.BodyParser(category); err != nil {
		db.Update(category.Name)
		return
	}
	db.Save(&category)
	c.Status(200).Send("Category Updated")

}
