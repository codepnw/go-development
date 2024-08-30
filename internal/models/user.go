package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id" gorm:"primaryKey;serializer:json;column:id"`
	FirstName string `json:"firstname" gorm:"serializer:json;column:firstname"`
	LastName  string `json:"lastname" gorm:"serializer:json;column:lastname"`
	Email     string `json:"email" gorm:"serializer:json;column:email"`
	CreatedAt string `json:"created_at" gorm:"serializer:json;column:created_at"`
}

func (u *User) Table() string { return "users" }

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
	return db.Table(u.Table()).Where("id = ?", &user.ID).Updates(&user).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Table(u.Table()).Unscoped().Delete(&u).Error
}

func (u *User) Get(db *gorm.DB, user User) *gorm.DB {
	return db.Table(u.Table()).Scopes(u.ValidateEmail(db, user))
}

func (u *User) GetAll(db *gorm.DB, user []User) *gorm.DB {
	return db.Table(u.Table())
}

func (u *User) GetById(db *gorm.DB) *gorm.DB {
	return db.Table(u.Table()).Where("").First(u)
}

func (u *User) GetByAttr(db *gorm.DB) *gorm.DB {
	return db.Table(u.Table()).First(u)
}

func (u *User) GetByIDs(ids []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(u.Table()).Where("id IN (?)", ids).First(u)
	}
}
