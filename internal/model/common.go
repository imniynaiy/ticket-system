package model

var DefaultPageReq = PageReq{
	Page: 0,
	Size: 10,
}

type PageReq struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
