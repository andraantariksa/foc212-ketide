package main

import (
	"github.com/go-xorm/xorm"

	"gitlab.com/cikadev/ketide/repository/codes"
	"gitlab.com/cikadev/ketide/repository/users"
)

func MigrateTable(db *xorm.Engine) {
	if err := db.Sync(new(users.Users)); err != nil {
		panic("Table users migrate error")
	}

	if err := db.Sync(new(codes.Codes)); err != nil {
		panic("Table code migrate error")
	}
}

// func dropTable(db *xorm.Engine) {
// 	if err := db.DropTables(new(Users)); err != nil {
// 		panic("Table users drop error")
// 	}

// 	if err := db.DropTables(new(Code)); err != nil {
// 		panic("Table code drop error")
// 	}
// }
