package models

import "time"

type Users struct {
	ID        int64     `xorm:"pk not null autoincr"`
	Username  string    `xorm:"unique not null"`
	Email     string    `xorm:"not null"`
	Password  string    `xorm:"not null"`
	CreatedAt time.Time `xorm:"created not null"`
	UpdatedAt time.Time `xorm:"updated not null"`
}
