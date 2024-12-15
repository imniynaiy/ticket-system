package service

import (
	"errors"

	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/model"
	"gorm.io/gorm"
)

func getSeatclassRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.Seatclass{})
}

func CreateSeatclass(req *model.CreateSeatclassReq) (*model.Seatclass, error) {
	seatclass := &model.Seatclass{
		SeatclassName: req.SeatclassName,
		Factor:        req.Factor,
	}

	if err := getSeatclassRepo().Create(seatclass).Error; err != nil {
		return nil, err
	}
	return seatclass, nil
}

func GetSeatclass(id uint) (*model.Seatclass, error) {
	var seatclass model.Seatclass
	if err := getSeatclassRepo().First(&seatclass, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("seatclass not found")
		}
		return nil, err
	}
	return &seatclass, nil
}

func UpdateSeatclass(req *model.UpdateSeatclassReq) (*model.Seatclass, error) {
	seatclass, err := GetSeatclass(req.SeatclassID)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if req.SeatclassName != "" {
		updates["seatclass_name"] = req.SeatclassName
	}
	if req.Factor != 0 {
		updates["factor"] = req.Factor
	}

	if err := getSeatclassRepo().Model(seatclass).Updates(updates).Error; err != nil {
		return nil, err
	}
	return seatclass, nil
}

func DeleteSeatclass(id uint) error {
	result := getSeatclassRepo().Delete(&model.Seatclass{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("seatclass not found")
	}
	return nil
}

func ListSeatclasses() (*model.ListSeatclassResp, error) {
	var total int64
	if err := getSeatclassRepo().Count(&total).Error; err != nil {
		return nil, err
	}

	var seatclasses []model.Seatclass
	if err := getSeatclassRepo().Find(&seatclasses).Error; err != nil {
		return nil, err
	}

	return &model.ListSeatclassResp{
		Total:       total,
		Seatclasses: seatclasses,
	}, nil
}
