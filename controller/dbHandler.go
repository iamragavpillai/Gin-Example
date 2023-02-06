package controller

import "gorm.io/gorm"

type handlerCustom struct {
	DB *gorm.DB
}

func NewConnection(db *gorm.DB) handlerCustom {
	return handlerCustom{db}
}
