package service

import (
	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/util"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func getUserRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.User{})
}

func Login(user *model.LoginReq) (token string, err error) {
	var userInDb model.User
	err = getUserRepo().Where("email = ?", user.Email).First(&userInDb).Error
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}
	err = util.CompareHashAndPassword([]byte(userInDb.PasswordHash), user.Password)
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}
	var us model.UserSession
	us.UserID = userInDb.UserId
	us.IsAdmin = userInDb.IsAdmin
	return util.GenTokenAndStoreInRedis(&us)
}

func Register(userReq *model.RegisterReq) error {
	// Check if email already exists
	var existingUser model.User
	err := getUserRepo().Where("email = ?", userReq.Email).First(&existingUser).Error
	if err == nil {
		return errors.ErrEmailAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		return errors.ErrInternalServerError
	}

	var newUser model.User
	copier.Copy(&newUser, userReq)
	hash, err := util.GeneratePasswordHash(userReq.Password)
	if err != nil {
		return errors.ErrInternalServerError
	}
	newUser.PasswordHash = hash
	err = getUserRepo().Create(&newUser).Error
	if err != nil {
		return errors.ErrInternalServerError
	}
	return nil
}
