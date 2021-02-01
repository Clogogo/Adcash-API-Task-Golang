package Models

import (
	"AdcashTask/Database"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"catName"`
}

func (category *Category) TableName() string {
	return "category"
}

//GetAllCategories Fetch all Category data
func GetAllCategories(category *[]Category) (err error) {
	if err = database.DB.Find(category).Error; err != nil {
		return err
	}
	return nil
}

//Get all product of a specific category
func CategoryProductList(id string) ([]*Product, error) {
	var products []*Product
	if err := database.DB.Where("cat_id = ?", id).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

//GetCategoryByID ... Fetch only one user by Id
func GetCategoryByID(category *Category, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(category).Error; err != nil {
		return err
	}
	return nil
}

//CreateCategory ... Insert New data
func CreateCategory(category *Category) (err error) {

	if err = database.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCategory ... Update category
func UpdateCategory(category *Category, id uint) (err error) {

	if err = database.DB.Where("id = ?", id).Save(category).Error; err != nil {
		return err
	}
	fmt.Println(category)
	return nil
}

//DeleteCategory ... Delete Category
func DeleteCategory(category *Category, id int64) (err error) {

	if err = database.DB.Where("id = ?", id).Delete(category).Error; err != nil {
		return err
	}
	return nil

}
