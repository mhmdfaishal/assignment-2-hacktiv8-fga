package models

import (
	"time"
)

type Order struct {
	OrderID      int       `gorm:"primary_key;auto_increment" json:"order_id"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}