package models

import (
	"time"

	"gorm.io/gorm"
)

type Movement struct {
	gorm.Model

	Id          string    `json:"id"`
	Date        time.Time `json:"time"`
	Account     string    `json:"account"`
	Amount      string    `json:"amount"`
	Description string    `json:"description"`
	Currency    string    `json:"currency"`
}
