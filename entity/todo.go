package entity

import "time"

type Todo struct {
	ID        int
	Title     string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
