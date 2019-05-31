package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type Customer struct {
	gorm.Model
	Name string
	Email string
}


func GetCustomers() ([]Customer, error) {

	db, err := GetConnection()
	var customers []Customer

	if err == nil {
		db.Find(&customers)
		defer db.Close()
	}

	return customers, err
}