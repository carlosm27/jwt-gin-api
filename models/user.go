package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type Repo struct {
	Db *gorm.DB
}

func RepoInterface(db *gorm.DB) *Repo {
	return &Repo{Db: db}
}

func (repo *Repo) CreateUser(user *User) (*User, error) {

	err := repo.Db.Create(&user).Error

	if err != nil {
		return &User{}, err
	}
	return user, nil
}
