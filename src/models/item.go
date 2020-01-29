package models

import (
	"time"
)

// Item defines properties of an item
type Item struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Price       float32    `json:"price"`
	CategoryID  int        `json:"category_id"`
	InSale      bool       `json:"in_sale"`
	AddedTime   *time.Time `json:"added_time"`
	LastUpdated *time.Time `json:"last_updated"`
	RemovedTime *time.Time `json:"removed_time"`
}
