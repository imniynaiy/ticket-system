package model

import "time"

type Route struct {
	RouteID       uint      `json:"route_id" gorm:"column:route_id;primaryKey;autoIncrement"`
	RouteName     string    `json:"route_name" gorm:"column:route_name"`
	DepartureTime time.Time `json:"departure_time" gorm:"column:departure_time"`
	ArrivalTime   time.Time `json:"arrival_time" gorm:"column:arrival_time"`
	DepartureFrom string    `json:"departure_from" gorm:"column:departure_from"`
	ArrivalTo     string    `json:"arrival_to" gorm:"column:arrival_to"`
	Distance      int       `json:"distance" gorm:"column:distance"`
	BasicFee      int       `json:"basic_fee" gorm:"column:basic_fee"`
}

type CreateRouteReq struct {
	RouteName     string    `json:"route_name" binding:"required"`
	DepartureTime time.Time `json:"departure_time" binding:"required"`
	ArrivalTime   time.Time `json:"arrival_time" binding:"required"`
	DepartureFrom string    `json:"departure_from" binding:"required"`
	ArrivalTo     string    `json:"arrival_to" binding:"required"`
	Distance      int       `json:"distance" binding:"required"`
	BasicFee      int       `json:"basic_fee" binding:"required"`
}

type UpdateRouteReq struct {
	RouteID       uint      `json:"route_id" binding:"required"`
	RouteName     string    `json:"route_name"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
	DepartureFrom string    `json:"departure_from"`
	ArrivalTo     string    `json:"arrival_to"`
	Distance      int       `json:"distance"`
	BasicFee      int       `json:"basic_fee"`
}

type ListRoutesReq struct {
	Page     int `form:"page,default=1" binding:"min=1"`
	PageSize int `form:"page_size,default=10" binding:"min=1,max=100"`
}

type ListRoutesResp struct {
	Total  int64   `json:"total"`
	Routes []Route `json:"routes"`
}

type UserListRoutesReq struct {
	Page          int       `form:"page,default=1" binding:"min=1"`
	PageSize      int       `form:"page_size,default=10" binding:"min=1,max=100"`
	DepartureFrom string    `form:"departure_from" binding:"required"`
	ArrivalTo     string    `form:"arrival_to" binding:"required"`
	DepartureTime time.Time `form:"departure_time" binding:"required"`
}

type RoutesWithSeatclass struct {
	Route
	SeatClass []ListRoutesSeatClass `json:"seat_class"`
}

type ListRoutesSeatClass struct {
	SeatClassID uint    `json:"seatclass_id" gorm:"column:seatclass_id"`
	SeatClass   string  `json:"seatclass_name" gorm:"column:seatclass_name"`
	Price       float32 `json:"price" gorm:"column:price"`
	Available   int     `json:"available" gorm:"column:available"`
}

type UserListRoutesResp struct {
	Total  int64                 `json:"total"`
	Routes []RoutesWithSeatclass `json:"routes"`
}

func (Route) TableName() string {
	return "route"
}
