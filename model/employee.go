package model

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `gorm:" not null" json:"email"`
	UserName      string    `gorm:"unique; not null; unique" json:"username"`
	EmailSentTime time.Time `json:"email_sent_time"`
	EmailSend     bool      `json:"email_send"`
}
