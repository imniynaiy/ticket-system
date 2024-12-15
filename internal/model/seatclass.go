package model

type Seatclass struct {
	SeatclassID   uint    `json:"seatclass_id" gorm:"column:seatclass_id;primaryKey;autoIncrement"`
	SeatclassName string  `json:"seatclass_name" gorm:"column:seatclass_name"`
	Factor        float64 `json:"factor" gorm:"column:factor"`
}

type CreateSeatclassReq struct {
	SeatclassName string  `json:"seatclass_name" binding:"required"`
	Factor        float64 `json:"factor" binding:"required"`
}

type UpdateSeatclassReq struct {
	SeatclassID   uint    `json:"seatclass_id" binding:"required"`
	SeatclassName string  `json:"seatclass_name"`
	Factor        float64 `json:"factor"`
}

type ListSeatclassResp struct {
	Total       int64       `json:"total"`
	Seatclasses []Seatclass `json:"seatclasses"`
}

func (Seatclass) TableName() string {
	return "seatclass"
}
