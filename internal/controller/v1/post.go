package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/log"
	"github.com/imniynaiy/ticket-system/internal/middleware"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

func AddPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")
	posts.GET("", getPosts)
	posts.POST("", middleware.Authenticationer(), addPost)
	posts.PUT("", middleware.Authenticationer(), modifyPost)
	posts.DELETE(":id", middleware.Authenticationer(), deletePost)
}

func getPosts(c *gin.Context) {
	var page model.PostPageReq
	var offset, limit int
	err := c.ShouldBindQuery(&page)
	if err != nil {
		log.Error("Failed to parse page, use default", log.String("err", err.Error()))
		offset = 0
		limit = 10
	} else {
		if page.Size <= 0 {
			limit = 10
		}
		if page.Page <= 0 {
			offset = 0
		} else {
			offset = (page.Page - 1) * page.Size
			limit = page.Size
		}
	}
	if page.Category == "all" {
		page.Category = ""
	}

	list, total, err := service.GetPostList(page.Category, offset, limit)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, model.PostPageResp{
		Posts: list,
		Count: int(total),
	})
}

func addPost(c *gin.Context) {
	var newPost model.Post
	err := c.ShouldBindBodyWithJSON(&newPost)
	if err != nil {
		log.Error("Failed to parse new post", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	err = service.AddPost(&newPost)
	if err != nil {
		log.Error("Failed to add new post", log.String("err", err.Error()))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func modifyPost(c *gin.Context) {
	var postToMod model.Post
	err := c.ShouldBindBodyWithJSON(&postToMod)
	if err != nil {
		log.Error("Failed to parse post", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	if postToMod.ID == 0 {
		log.Error("Empty post id")
		c.Status(http.StatusBadRequest)
		return
	}
	err = service.ModifyPost(&postToMod)
	if err != nil {
		log.Error("Failed to update new post", log.Uint("id", postToMod.ID), log.String("err", err.Error()))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func deletePost(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		log.Error("Failed to parse id", log.String("id", idString))
		c.Status(http.StatusBadRequest)
		return
	}
	err = service.DeletePost(uint(id))
	if err != nil {
		log.Error("Failed to delete post", log.String("id", idString), log.String("err", err.Error()))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
