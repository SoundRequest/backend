package structure

import (
	"time"
)

// User Database structure
type User struct {
	ID                 int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Name               string `gorm:"type:varchar(50);unique_index;NOT NULL"`
	Password           string `gorm:"type:varchar(70);unique_index;NOT NULL"`
	Email              string `gorm:"type:varchar(60);unique_index;NOT NULL"`
	Verified           bool   `gorm:"NOT NULL;defualt:default:0"`
	VerifyCode         string `gorm:"type:varchar(10);unique_index;NOT NULL"`
	VerifyCodePassword string `gorm:"type:varchar(10);"`
	LastPurchasedDate  *time.Time
}

type UserSetting struct {
	Author User   `gorm:"not null;foreignkey:ID"`
	Dark   bool   `gorm:"NOT NULL;defualt:default:0"`
	Lang   string `gorm:"NOT NULL;defualt:default:ko"`
}

type PlayItem struct {
	ID          int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Author      User   `gorm:"not null;foreignkey:ID"`
	Name        string `gorm:"type:varchar(50);unique_index;NOT NULL"`
	Description string `gorm:"type:varchar(1000);unique_index;NOT NULL"`
	Link        string `gorm:"type:varchar(500);NOT NULL"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type PlayList struct {
	ID          int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Author      User   `gorm:"not null;foreignkey:ID"`
	Name        string `gorm:"type:varchar(50);unique_index;NOT NULL"`
	Description string `gorm:"type:varchar(1000);unique_index;NOT NULL"`
	Public      bool   `gorm:"NOT NULL;defualt:default:0"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type PlayTag struct {
	ID        int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Author    User   `gorm:"not null;foreignkey:ID"`
	Name      string `gorm:"type:varchar(50);unique_index;NOT NULL"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type playBridge struct {
	Item     int `gorm:"NOT NULL"`
	PlayList int
	PlayTag  int
}
