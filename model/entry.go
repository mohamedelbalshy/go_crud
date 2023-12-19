package model

import (
	"CRUD/api/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	ID      int64  `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Content string `gorm:"type:text" json:"content"`
	UserID  int64
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
