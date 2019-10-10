package models

import "time"

type Codes struct {
	ID        int64  `xorm:"pk not null autoincr"`
	Language  string `xorm:"not null"`
	Code      string `xorm:"not null"`
	Stdin     string
	Stdout    string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
