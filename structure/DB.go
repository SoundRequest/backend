package structure

import (
	"time"
)

// Authentication Database structure
type Authentication struct {
	ID                int64     `gorm:"not null;AUTO_INCREMENT;unique_index"`
	ClientID          string    `gorm:"type:varchar(45)"`
	ClientSecret      string    `gorm:"type:varchar(45)"`
	ClientToken       string    `gorm:"type:varchar(45)"`
	Username          string    `gorm:"type:varchar(50)"`
	PasswordHash      string    `gorm:"type:varchar(50)"`
	Email             string    `gorm:"type:varchar(60)"`
	LastPurchasedDate time.Time `gorm:"type:timestamp"`
}
