package models

import (
	"time"
)

type Name struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type Update_user struct {
	Id    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

type Users struct {
	Id        int       `json:"id"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Email     string    `json:"email"`
	Pwd       string    `json:"pwd"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
