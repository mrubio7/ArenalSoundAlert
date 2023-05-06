package db

import (
	"ArenalSoundAlert/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB = nil

func Connection() *gorm.DB {
	if Db != nil {
		return Db
	}

	connectionString := "host=192.168.1.29 user=postgres password=**** dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	newDb, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("No se puede conectar a la bd")
	}

	return newDb
}

func NuevoElemento(media models.Media) {
	db := Connection()

	db.Create(&media)
	db.Commit()
}

func GetUltimoElemento() models.Media {
	var media models.Media
	db := Connection()

	db.Order("date desc").Last(&media)
	return media
}
