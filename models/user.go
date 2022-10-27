package models

import (
	"errors"
	"github.com/carlosm27/jwtGinApi/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"log"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	var err error

	user := User{}

	db, err := Setup()
	if err != nil {
		log.Println(err)
		return "", err
	}

	if err = db.Model(User{}).Where("username=?", username).Take(&user).Error; err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func GetUserByID(uid uint) (User, error) {
	var user User

	db, err := Setup()

	if err != nil {
		log.Println(err)
		return User{}, err
	}
	if err := db.First(&user, uid).Error; err != nil {
		return user, errors.New("user not found")

	}
	user.Password = ""

	return user, nil
}
