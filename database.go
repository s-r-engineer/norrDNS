package main

import (
	libraryErrors "github.com/s-r-engineer/library/errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define User model
type queryHistory struct {
	gorm.Model
	QueriedAddress  string
	ResolvedAddress string
	ResolvedLoad    int
	CountryId       int
	CountryCode     string
}

var db *gorm.DB

func initDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("/database/queries.db"), &gorm.Config{})
	libraryErrors.Panicer(err)
	err = db.AutoMigrate(&queryHistory{})
	libraryErrors.Panicer(err)
}

func add(q queryHistory) {
	if db != nil {
		libraryErrors.Errorer(db.Create(&q).Error)
	}
}

func getLatestForTheCountryInLastOneMinute(countryId int) (queryHistory, error) {
	server := queryHistory{}
	return server, db.Where("country_id = ?", countryId).Order("created_at desc").First(&server).Error
}
