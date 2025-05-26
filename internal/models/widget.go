package models

import "time"

type Widget struct {
	ID          int
	Name        string
	Description string

	Components []WidgetComponent

	CreatedAt time.Time
	UpdatedAt time.Time
}

type WidgetComponent struct {
	ID         int
	Name       string
	WidgetID   int
	Complexity int

	CreatedAt string
	UpdatedAt string
}
