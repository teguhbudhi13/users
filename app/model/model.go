package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User schema
type User struct {
	ID   int       `gorm:"primary_key; auto_increment" sql:"id" json:"id"`
	Name string    `gorm:"type:varchar; not null" json:"name"`
	DOB  time.Time `gorm:"type:date; not null" json:"dob"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
