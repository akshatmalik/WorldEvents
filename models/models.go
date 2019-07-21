package models

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type BaseEvent struct {
	gorm.Model
	Title         string
	EventLocation Location
	EventSummary  Summary
	EventTime     time.Time
}

type Summary struct {
	Introduction string
	Details      map[string]string
}

type Location struct {
	gorm.Model
	Continent string
	Country   string
	City      string
	State     string
	Address   string
}

type TextEvent struct {
	BaseEvent
	TextDescription string
}

type Event interface {
	GetDate()
	GetName()
}
