package entities

import "time"

type Api struct {
	Id        uint
	Name      string
	Url       string
	Method    string
	Response  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
