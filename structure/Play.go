package structure

import (
	"time"
)

type PlayItem struct {
	ID          int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Author      User   `gorm:"not null;foreignkey:ID"`
	Name        string `gorm:"type:varchar(50);NOT NULL"`
	Description string `gorm:"type:varchar(1000);NOT NULL"`
	Link        string `gorm:"type:varchar(500);NOT NULL"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type PlayList struct {
	ID          int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Author      User   `gorm:"not null;foreignkey:ID"`
	Name        string `gorm:"type:varchar(50);NOT NULL"`
	Description string `gorm:"type:varchar(1000);NOT NULL"`
	Public      bool   `gorm:"NOT NULL;defualt:default:0"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type PlayTag struct {
	ID        int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Author    User   `gorm:"not null;foreignkey:ID"`
	Name      string `gorm:"type:varchar(50);NOT NULL"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type PlayBridge struct {
	Item     int `gorm:"NOT NULL"`
	PlayList int
	PlayTag  int
}
