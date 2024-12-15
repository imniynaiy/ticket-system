package service

import (
	"errors"

	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/model"
	"gorm.io/gorm"
)

func getRouteRepo() *gorm.DB {
	return database.GlobalDB.Model(&model.Route{})
}

func CreateRoute(req *model.CreateRouteReq) (*model.Route, error) {
	route := &model.Route{
		RouteName:     req.RouteName,
		DepartureTime: req.DepartureTime,
		ArrivalTime:   req.ArrivalTime,
		DepartureFrom: req.DepartureFrom,
		ArrivalTo:     req.ArrivalTo,
		Distance:      req.Distance,
		BasicFee:      req.BasicFee,
	}

	if err := getRouteRepo().Create(route).Error; err != nil {
		return nil, err
	}
	return route, nil
}

func GetRoute(routeID uint) (*model.Route, error) {
	var route model.Route
	if err := getRouteRepo().First(&route, routeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("route not found")
		}
		return nil, err
	}
	return &route, nil
}

func UpdateRoute(req *model.UpdateRouteReq) (*model.Route, error) {
	route, err := GetRoute(req.RouteID)
	if err != nil {
		return nil, err
	}

	// Only update non-zero values
	updates := map[string]interface{}{}
	if req.RouteName != "" {
		updates["route_name"] = req.RouteName
	}
	if !req.DepartureTime.IsZero() {
		updates["departure_time"] = req.DepartureTime
	}
	if !req.ArrivalTime.IsZero() {
		updates["arrival_time"] = req.ArrivalTime
	}
	if req.DepartureFrom != "" {
		updates["departure_from"] = req.DepartureFrom
	}
	if req.ArrivalTo != "" {
		updates["arrival_to"] = req.ArrivalTo
	}
	if req.Distance != 0 {
		updates["distance"] = req.Distance
	}
	if req.BasicFee != 0 {
		updates["basic_fee"] = req.BasicFee
	}

	if err := getRouteRepo().Model(route).Updates(updates).Error; err != nil {
		return nil, err
	}
	return route, nil
}

func DeleteRoute(routeID uint) error {
	result := getRouteRepo().Delete(&model.Route{}, routeID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("route not found")
	}
	return nil
}

func ListRoutes(page, pageSize int) (*model.ListRoutesResp, error) {
	var total int64
	if err := getRouteRepo().Model(&model.Route{}).Count(&total).Error; err != nil {
		return nil, err
	}

	var routes []model.Route
	offset := (page - 1) * pageSize
	if err := getRouteRepo().Offset(offset).Limit(pageSize).Find(&routes).Error; err != nil {
		return nil, err
	}

	return &model.ListRoutesResp{
		Total:  total,
		Routes: routes,
	}, nil
}
