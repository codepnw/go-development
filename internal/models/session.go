package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	ID string `gorm:"unique"`
	Data []byte
	Expiry int64
}

func SessionFile() string {
	return "session.db"
}

