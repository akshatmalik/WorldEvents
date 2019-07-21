package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type BaseEvent struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Title     string
}

type TextEvent struct {
	BaseEvent
	TextDescription string
}

type Event interface {
	GetDate()
	GetName()
}
