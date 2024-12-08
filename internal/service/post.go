package service

import (
	"errors"

	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/model"
	"gorm.io/gorm"
)

func getPostRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.Post{})
}
func GetPostList(category string, offset int, limit int) (list []model.Post, total int64, err error) {
	db := getPostRepo()
	if category != "" {
		db = db.Where("category = ?", category)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var postList []model.Post
	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&postList).Error
	return postList, total, err
}

func AddPost(newPost *model.Post) error {
	return getPostRepo().Create(newPost).Error
}

func ModifyPost(postToMod *model.Post) error {
	result := getPostRepo().Where("ID = ?", postToMod.ID).Updates(postToMod)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func DeletePost(id uint) error {
	db := database.GlobalDB.Model(&model.Post{})
	return db.Delete(&model.Post{}, id).Error
}
