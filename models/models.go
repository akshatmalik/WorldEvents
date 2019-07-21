package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type BaseEvent struct {
	ID        uint `gorm: "primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Tit
}

type TextEvent struct {
	gorm.BaseEvent
	BaseEvent
}

type Event interface {
	GetDate()
	GetName()
}
