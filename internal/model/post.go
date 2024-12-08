package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content  string
	Category string
}

func (Post) TableName() string {
	return "t_posts"
}

type PostPageReq struct {
	PageReq
	Category string `form:"category"`
}

type PostPageResp struct {
	Posts []Post
	Count int
}
