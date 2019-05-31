package repository

import (
	"environment.go.example/app/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
)


var instance *gorm.DB
var err error
var once sync.Once

func GetConnection() (*gorm.DB, error) {
	once.Do(func() {
		instance, err = gorm.Open("postgres", config.GetSource())

		instance.LogMode(!config.IsProduction())
	})
	return instance, err
}