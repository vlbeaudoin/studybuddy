package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Note struct {
	gorm.Model
	Word       string `json:"word"`
	Definition string `json:"definition"`
	Category   string `json:"category"`
}

type ShortNote struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
	Category   string `json:"category"`
}

func OpenDatabase() error {
	var err error

	db, err = gorm.Open(sqlite.Open("studybuddy.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func MigrateDatabase() {
	db.AutoMigrate(&Note{})
}

func InsertNote(word string, definition string, category string) {
	db.Create(&Note{
		Word:       word,
		Definition: definition,
		Category:   category,
	})
}

func ListNotes() ([]ShortNote, error) {
	var shortNotes []ShortNote

	result := db.Model(&Note{}).Find(&shortNotes)

	return shortNotes, result.Error
}
