package model

type Seat struct {
	RouteID     uint `json:"route_id" gorm:"column:route_id"`
	SeatID      uint `json:"seat_id" gorm:"column:seat_id;primaryKey"`
	SeatclassID uint `json:"seatclass_id" gorm:"column:seatclass_id"`
	Status      uint `json:"status" gorm:"column:status"`
}

type CreateSeatReq struct {
	RouteID     uint `json:"route_id" binding:"required"`
	SeatclassID uint `json:"seatclass_id" binding:"required"`
	Status      uint `json:"status" binding:"required"`
}

type UpdateSeatReq struct {
	RouteID     uint `json:"route_id"`
	SeatID      uint `json:"seat_id" binding:"required"`
	SeatclassID uint `json:"seatclass_id"`
	Status      uint `json:"status"`
}

type ListSeatsReq struct {
	RouteID  uint `form:"route_id"`
	Page     int  `form:"page,default=1" binding:"min=1"`
	PageSize int  `form:"page_size,default=10" binding:"min=1,max=100"`
}

type ListSeatsResp struct {
	Total int64  `json:"total"`
	Seats []Seat `json:"seats"`
}

type UserListSeatsReq struct {
	RouteID  uint `form:"route_id" binding:"required"`
	Page     int  `form:"page,default=1" binding:"min=1"`
	PageSize int  `form:"page_size,default=10" binding:"min=1,max=100"`
}

type UserListSeatsResp struct {
	Total int64  `json:"total"`
	Seats []Seat `json:"seats"`
}

func (Seat) TableName() string {
	return "seat"
}
