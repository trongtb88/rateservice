package models

import "time"

type CurrencyRate struct {
	Date    time.Time      `gorm:"primaryKey;size:15;not null" json:"date"`
	CurrencyCode  string   `gorm:"primaryKey;size:5;not null" json:"currency_code"`
	Rate     float64       `gorm:"not null;type:decimal(15,6)" json:"rate"`
	CreatedAt time.Time    `gorm:""DEFAULT:current_timestamp; type:timestamp"" json:"created_at"`
	UpdatedAt time.Time    `gorm:""DEFAULT:current_timestamp;type:timestamp"" json:"updated_at"`
}
