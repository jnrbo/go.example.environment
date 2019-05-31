package repository

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type Customer struct {
	ID uint `gorm:"primary_key"`
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