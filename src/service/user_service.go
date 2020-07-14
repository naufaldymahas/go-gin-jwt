package service

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/naufaldymahas/go-gin-jwt/src/model"
	"github.com/naufaldymahas/go-gin-jwt/src/payload"
	"github.com/naufaldymahas/go-gin-jwt/src/util"
)

type userService struct {
	DB *gorm.DB
}

func UserService(db *gorm.DB) userService {
	return userService{
		DB: db,
	}
}

func (us *userService) CreateUser(auth payload.AuthRegister) error {
	var user model.User

	user.Name = auth.Name
	user.Email = auth.Email
	user.Password = util.HashPassword(auth.Password)

	tx := us.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (us *userService) LoginUser(auth payload.AuthLogin) (string, bool) {
	var user model.User

	if err := us.DB.Where("email = ?", auth.Email).First(&user).Error; err != nil {
		log.Println(err.Error())
		return "", false
	}

	if user.Password != util.HashPassword(auth.Password) {
		log.Println("Authentication Failed!")
		return "", false
	}

	token, err := util.TokenGenerator(user.Email)
	if err != nil {
		log.Println("JWT token failed:", token)
		return "", false
	}

	return token, true
}

func (us *userService) FindByEmail(email string) (model.User, error) {
	var user model.User

	if err := us.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}

	return user, nil
}
