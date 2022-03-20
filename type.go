package main

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Post   string `gorm:"type: varchar(300)" json:"post" binding:"required"`
	Desc   string `gorm:"type: varchar(500)" json:"desc" binding:"required"`
	Status string
}
