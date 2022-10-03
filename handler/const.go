package handler

import "reflect"

var ConfigTypeMap = map[string]reflect.Type{
	"personal_check_handler": reflect.TypeOf(PersonalCheckHandler{}),
	"city_check_handler":     reflect.TypeOf(CityCheckHandler{}),
	"new_check_handler":      reflect.TypeOf(NewCheckHandler{}),
	"recent_check_handler":   reflect.TypeOf(RecentCheckHandler{}),
}
