package service

import (
	"errors"
	"time"

	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getReservationRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.Reservation{})
}

func getReservationDetailRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.ReservationDetail{})
}

func GetUserReservationWithDetails(id uint, userId uint) (*model.ReservationWithDetail, error) {
	var reservation model.Reservation
	if err := getReservationRepo().Where("reservation_id = ? AND user_id = ?", id, userId).First(&reservation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation not found")
		}
		return nil, err
	}

	var detail model.ReservationDetail
	if err := getReservationDetailRepo().Where("reservation_id = ?", id).First(&detail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation detail not found")
		}
		return nil, err
	}

	return &model.ReservationWithDetail{
		Reservation:         reservation,
		ReservationDetailID: detail.ReservationDetailID,
		RouteID:             detail.RouteID,
		SeatID:              detail.SeatID,
		PassengerFamilyName: detail.PassengerFamilyName,
		PassengerFirstName:  detail.PassengerFirstName,
	}, nil
}

func GetReservationWithDetails(id uint) (*model.ReservationWithDetail, error) {
	var reservation model.Reservation
	if err := getReservationRepo().First(&reservation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation not found")
		}
		return nil, err
	}

	var detail model.ReservationDetail
	if err := getReservationDetailRepo().Where("reservation_id = ?", id).First(&detail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation detail not found")
		}
		return nil, err
	}

	return &model.ReservationWithDetail{
		Reservation:         reservation,
		ReservationDetailID: detail.ReservationDetailID,
		RouteID:             detail.RouteID,
		SeatID:              detail.SeatID,
		PassengerFamilyName: detail.PassengerFamilyName,
		PassengerFirstName:  detail.PassengerFirstName,
	}, nil
}

func ListUserReservations(userId uint, req *model.ListReservationsReq) (*model.ListReservationsResp, error) {
	db := getReservationRepo().Where("user_id = ?", userId)

	if req.RouteID != 0 {
		db = db.Where("route_id = ?", req.RouteID)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if !req.StartTime.IsZero() {
		db = db.Where("reservation_date >= ?", req.StartTime.Format("2006-01-02"))
	}
	if !req.EndTime.IsZero() {
		endTime := req.EndTime.Add(24 * time.Hour) // Include the entire end day
		db = db.Where("reservation_date < ?", endTime.Format("2006-01-02"))
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	var reservations []model.Reservation
	offset := (req.Page - 1) * req.PageSize
	if err := db.Offset(offset).Limit(req.PageSize).
		Find(&reservations).Error; err != nil {
		return nil, err
	}

	// Get details for all reservations
	result := make([]model.ReservationWithDetail, 0, len(reservations))
	for _, res := range reservations {
		var detail model.ReservationDetail
		if err := getReservationDetailRepo().
			Where("reservation_id = ?", res.ReservationID).
			First(&detail).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
			// Skip if detail not found
			continue
		}

		result = append(result, model.ReservationWithDetail{
			Reservation:         res,
			ReservationDetailID: detail.ReservationDetailID,
			RouteID:             detail.RouteID,
			SeatID:              detail.SeatID,
			PassengerFamilyName: detail.PassengerFamilyName,
			PassengerFirstName:  detail.PassengerFirstName,
		})
	}

	return &model.ListReservationsResp{
		Total:        total,
		Reservations: result,
	}, nil
}

func ListReservations(req *model.ListReservationsReq) (*model.ListReservationsResp, error) {
	db := getReservationRepo()

	if req.RouteID != 0 {
		db = db.Where("route_id = ?", req.RouteID)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if !req.StartTime.IsZero() {
		db = db.Where("reservation_date >= ?", req.StartTime.Format("2006-01-02"))
	}
	if !req.EndTime.IsZero() {
		endTime := req.EndTime.Add(24 * time.Hour) // Include the entire end day
		db = db.Where("reservation_date < ?", endTime.Format("2006-01-02"))
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	var reservations []model.Reservation
	offset := (req.Page - 1) * req.PageSize
	if err := db.Offset(offset).Limit(req.PageSize).
		Find(&reservations).Error; err != nil {
		return nil, err
	}

	// Get details for all reservations
	result := make([]model.ReservationWithDetail, 0, len(reservations))
	for _, res := range reservations {
		var detail model.ReservationDetail
		if err := getReservationDetailRepo().
			Where("reservation_id = ?", res.ReservationID).
			First(&detail).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
			// Skip if detail not found
			continue
		}

		result = append(result, model.ReservationWithDetail{
			Reservation:         res,
			ReservationDetailID: detail.ReservationDetailID,
			RouteID:             detail.RouteID,
			SeatID:              detail.SeatID,
			PassengerFamilyName: detail.PassengerFamilyName,
			PassengerFirstName:  detail.PassengerFirstName,
		})
	}

	return &model.ListReservationsResp{
		Total:        total,
		Reservations: result,
	}, nil
}

func CreateReservation(userId uint, req *model.CreateReservationReq) (*model.ReservationWithDetail, error) {
	db := database.GlobalDB

	// Create initial log entry
	log := &model.Log{
		UserID:     userId,
		FunctionID: 1,
		Timestamp:  time.Now(),
		Result:     0,
	}
	if err := getLogRepo().Create(log).Error; err != nil {
		return nil, err
	}

	// Start transaction
	tx := db.Begin()
	if tx.Error != nil {
		getLogRepo().Model(log).Updates(map[string]interface{}{
			"result": 2,
		})
		return nil, tx.Error
	}

	// Create reservation
	reservation := &model.Reservation{
		UserID:          userId,
		ReservationDate: req.ReservationDate.Format("2006-01-02"),
	}

	if err := tx.Create(reservation).Error; err != nil {
		tx.Rollback()
		getLogRepo().Model(log).Updates(map[string]interface{}{
			"result": 2,
		})
		return nil, err
	}

	// Create reservation detail
	detail := &model.ReservationDetail{
		ReservationID:       reservation.ReservationID,
		RouteID:             req.RouteID,
		SeatID:              req.SeatID,
		PassengerFamilyName: req.PassengerFamilyName,
		PassengerFirstName:  req.PassengerFirstName,
	}

	// Get route_id from seat table
	var seat model.Seat
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("seat_id = ? and status = 1", req.SeatID).First(&seat).Error; err != nil {
		tx.Rollback()
		getLogRepo().Model(log).Updates(map[string]interface{}{
			"result": 2,
		})
		return nil, err
	}
	// Update route_id from seat
	detail.RouteID = seat.RouteID
	if err := tx.Create(detail).Error; err != nil {
		tx.Rollback()
		getLogRepo().Model(log).Updates(map[string]interface{}{
			"result": 2,
		})
		return nil, err
	}

	// Update seat status
	if err := tx.Model(&model.Seat{}).Where("seat_id = ? and status = 1", req.SeatID).
		Update("status", 0).Error; err != nil {
		tx.Rollback()
		getLogRepo().Model(log).Updates(map[string]interface{}{
			"result": 2,
		})
		return nil, err
	}

	// Update log status to success
	if err := getLogRepo().Model(log).Update("result", 1).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		getLogRepo().Model(log).Updates(map[string]interface{}{
			"result": 2,
		})
		return nil, err
	}

	return &model.ReservationWithDetail{
		Reservation:         *reservation,
		ReservationDetailID: detail.ReservationDetailID,
		RouteID:             detail.RouteID,
		SeatID:              detail.SeatID,
		PassengerFamilyName: detail.PassengerFamilyName,
		PassengerFirstName:  detail.PassengerFirstName,
	}, nil
}
