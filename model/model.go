package model

import (
	"time"
)

type ExchangeRate struct {
	From string `json:"from" gorm:"primary_key"`
	To string 	`json:"to" gorm:"primary_key"`
	Date time.Time  `json:"date" gorm:"primary_key"`
	Rate float64 `json:"rate"`
}

type TrackedExchange struct {
	User string `json:"user" gorm:"primary_key"`
	From string `json:"from" gorm:"primary_key"`
	To string 	`json:"to" gorm:"primary_key"`
}