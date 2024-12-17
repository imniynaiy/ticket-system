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

func UserListRoutes(req *model.UserListRoutesReq) (*model.UserListRoutesResp, error) {
	var total int64
	query := getRouteRepo().Model(&model.Route{}).
		Where("departure_from = ? AND arrival_to = ? AND DATE(departure_time) = DATE(?)",
			req.DepartureFrom, req.ArrivalTo, req.DepartureTime)

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var routes []model.Route
	offset := (req.Page - 1) * req.PageSize
	if err := query.Offset(offset).Limit(req.PageSize).Find(&routes).Error; err != nil {
		return nil, err
	}
	// Get seat class information for each route
	routesWithSeats := make([]model.RoutesWithSeatclass, len(routes))
	for i, route := range routes {
		// Copy base route info
		routesWithSeats[i].Route = route

		// Get seat classes and counts for this route
		var seatClasses []model.ListRoutesSeatClass
		err := database.GlobalDB.Raw(`
		SELECT 
			sc.seatclass_id,
			sc.seatclass_name,
			sc.factor * r.basic_fee as price,
			COUNT(*) as available
		FROM seatclass sc
		JOIN seat s ON s.seatclass_id = sc.seatclass_id 
		JOIN route r ON s.route_id = r.route_id
		WHERE r.route_id = ? and s.status = 1
		GROUP BY sc.seatclass_id, sc.seatclass_name, sc.factor, r.basic_fee;
		`, route.RouteID).Scan(&seatClasses).Error

		if err != nil {
			return nil, err
		}

		routesWithSeats[i].SeatClass = seatClasses
	}

	return &model.UserListRoutesResp{
		Total:  total,
		Routes: routesWithSeats,
	}, nil
}
