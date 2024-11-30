package utils

import (
	"reflect"
	"strings"
)

const (
	SUCCESS = 200
)

const (
	ERR_SUCCESS = 1
	ERR_FAILED  = 0
	ERR_ERROR   = -1
)

const (
	IDLE_IDENTITY   = 0
	IDLE_TYPEENTITY = -1
	RESERVED_PRICE  = -1.00
)

const (
	NIL_STRING = "\xff"
)

func Success(errcode interface{}) bool {
	val := reflect.ValueOf(errcode).Int()
	return val == ERR_SUCCESS
}

func Failed(errcode interface{}) bool {
	val := reflect.ValueOf(errcode).Int()
	return val == ERR_FAILED
}

func Error(errcode interface{}) bool {
	val := reflect.ValueOf(errcode).Int()
	return val != ERR_SUCCESS
}

func GetIdleId() int64 {
	return IDLE_IDENTITY
}

func IsValidId(id int64) bool {
	return id != IDLE_IDENTITY
}

func IsValidType(typeValue int32) bool {
	return typeValue != IDLE_TYPEENTITY
}

func IsValidField(field string) bool {
	field2 := strings.TrimSpace(field)
	return field2 != ""
}

func IsValidPrice(price float64) bool {
	return price >= 0.0 || price == RESERVED_PRICE
}

func IsValidPage(pageNum, pageSize int32) bool {
	return pageNum >= 1 && pageSize >= 1
}

func IsValidTime(time int64) bool {
	return time > 0
}
