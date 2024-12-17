package model

import "time"

type Reservation struct {
	ReservationID   uint   `json:"reservation_id" gorm:"column:reservation_id;primaryKey;autoIncrement"`
	UserID          uint   `json:"user_id" gorm:"column:user_id"`
	ReservationDate string `json:"reservation_date" gorm:"column:reservation_date"`
}

type ReservationDetail struct {
	ReservationID       uint   `json:"reservation_id" gorm:"column:reservation_id;primaryKey"`
	ReservationDetailID uint   `json:"reservation_detail_id" gorm:"column:reservation_detail_id"`
	RouteID             uint   `json:"route_id" gorm:"column:route_id"`
	SeatID              uint   `json:"seat_id" gorm:"column:seat_id"`
	PassengerFamilyName string `json:"passenger_family_name" gorm:"column:passenger_family_name"`
	PassengerFirstName  string `json:"passenger_first_name" gorm:"column:passenger_first_name"`
}

type ReservationWithDetail struct {
	Reservation
	ReservationDetailID uint   `json:"reservation_detail_id"`
	RouteID             uint   `json:"route_id"`
	SeatID              uint   `json:"seat_id"`
	PassengerFamilyName string `json:"passenger_family_name"`
	PassengerFirstName  string `json:"passenger_first_name"`
}

type ListReservationsReq struct {
	RouteID   uint      `form:"route_id"`
	Status    *uint     `form:"status"`
	StartTime time.Time `form:"start_time" time_format:"2006-01-02"`
	EndTime   time.Time `form:"end_time" time_format:"2006-01-02"`
	Page      int       `form:"page,default=1" binding:"min=1"`
	PageSize  int       `form:"page_size,default=10" binding:"min=1,max=100"`
}

type ListReservationsResp struct {
	Total        int64                   `json:"total"`
	Reservations []ReservationWithDetail `json:"reservations"`
}

func (Reservation) TableName() string {
	return "reservation"
}

func (ReservationDetail) TableName() string {
	return "reservation_detail"
}

type CreateReservationReq struct {
	RouteID             uint      `json:"route_id"`
	SeatID              uint      `json:"seat_id"`
	PassengerFamilyName string    `json:"passenger_family_name"`
	PassengerFirstName  string    `json:"passenger_first_name"`
	ReservationDate     time.Time `json:"reservation_date"`
}
