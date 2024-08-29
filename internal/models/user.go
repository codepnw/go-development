package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id" gorm:"primaryKey;column:user_id;serializer:json"`
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) Table() string { return "dbo.users" }

func (u *User) ValidateEmail(db *gorm.DB, user User) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if user.Email != "" {
			return db
		}
		db.AddError(errors.New("invalid email provided"))
		return db
	}
}

func (u *User) Create(db *gorm.DB) error {
	if err := db.Table(u.Table()).Scopes(u.ValidateEmail(db, *u)).Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Update(db *gorm.DB, user User) error {
	return db.Table(u.Table()).Updates(&u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Table(u.Table()).Unscoped().Delete(&u).Error
}

func (u *User) Get(db *gorm.DB) *gorm.DB {
	return db.Table(u.Table())
}

func (u *User) GetByEmail(db *gorm.DB, email string) *gorm.DB {
	return db.Table(u.Table()).Where("email = ?", `"`+email+`"`).First(u)
}

func (u *User) GetByID() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(u.Table()).Where("id = ?", u.ID).First(u)
	}
}

func (u *User) GetByIDs(ids []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(u.Table()).Where("id IN (?)", ids).First(u)
	}
}
