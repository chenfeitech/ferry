// +build sqlite3

package database

import (
	_ "gorm.io/driver/sqlite"
)

type SqlLite struct {
}

func (*SqlLite) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}
