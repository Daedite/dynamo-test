package models

import "time"

type ShopperProfile struct {
	ID                 uint   `fake:"{number:0,10000}"`
	Email              string `fake:"{email}"`
	Currency           string
	OutstandingBalance float64 `fake:"{number:0,10000}"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
type ShopperProfileSample struct {
	ID        uint   `fake:"{number:0,10000}"`
	Email     string `fake:"{email}"`
	Currency  string
	Balance   float64 `fake:"{number:0,10000}"`
	CreatedAt string
	UpdatedAt string
}
