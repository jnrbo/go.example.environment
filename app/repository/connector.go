package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.example.environment/app/config"
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