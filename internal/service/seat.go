package service

import (
	"errors"

	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/model"
	"gorm.io/gorm"
)

func getSeatRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.Seat{})
}

func CreateSeat(req *model.CreateSeatReq) (*model.Seat, error) {
	seat := &model.Seat{
		RouteID:     req.RouteID,
		SeatclassID: req.SeatclassID,
		Status:      req.Status,
	}

	if err := getSeatRepo().Create(seat).Error; err != nil {
		return nil, err
	}
	return seat, nil
}

func GetSeat(seatID uint) (*model.Seat, error) {
	var seat model.Seat
	if err := getSeatRepo().First(&seat, "seat_id = ?", seatID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("seat not found")
		}
		return nil, err
	}
	return &seat, nil
}

func UpdateSeat(req *model.UpdateSeatReq) (*model.Seat, error) {
	seat, err := GetSeat(req.SeatID)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if req.RouteID != 0 {
		updates["route_id"] = req.RouteID
	}
	if req.SeatclassID != 0 {
		updates["seatclass_id"] = req.SeatclassID
	}

	updates["status"] = req.Status

	if err := getSeatRepo().Model(seat).Updates(updates).Error; err != nil {
		return nil, err
	}
	return seat, nil
}

func DeleteSeat(seatID uint) error {
	result := getSeatRepo().Delete(&model.Seat{}, "seat_id = ?", seatID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("seat not found")
	}
	return nil
}

func ListSeats(req *model.ListSeatsReq) (*model.ListSeatsResp, error) {
	db := getSeatRepo()
	if req.RouteID != 0 {
		db = db.Where("route_id = ?", req.RouteID)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	var seats []model.Seat
	offset := (req.Page - 1) * req.PageSize
	if err := db.Offset(offset).Limit(req.PageSize).Find(&seats).Error; err != nil {
		return nil, err
	}

	return &model.ListSeatsResp{
		Total: total,
		Seats: seats,
	}, nil
}
