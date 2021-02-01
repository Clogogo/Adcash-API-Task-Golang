package Models

import (
	database "AdcashTask/Database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"product_Name"`
	CatId int    `json:"category_Id"`
}

//GetAllProduct Fetch all Product data
func GetAllProduct(product *[]Product) (err error) {

	if err = database.DB.Find(product).Error; err != nil {
		return err
	}
	return nil

}

//CreateProduct ... Insert New data
func CreateProduct(product *Product) (err error) {

	if err = database.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetProductByID ... Fetch only one user by Id
func GetProductByID(product *Product, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProduct ... Update Product
func UpdateProduct(product *Product, id uint) (err error) {

	if err = database.DB.Where("id = ?", id).Save(product).Error; err != nil {
		return err
	}
	fmt.Println(product)
	return nil

}

//DeleteProduct ... Delete Product
func DeleteProduct(product *Product, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(product)
	return nil
}
