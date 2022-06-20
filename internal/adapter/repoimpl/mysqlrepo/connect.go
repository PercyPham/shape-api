package mysqlrepo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(connectLink string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(connectLink), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open gorm db connection: %w", err)
	}
	return db, nil
}
