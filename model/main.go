package model

import "github.com/kilgaloon/atm/utils"

// Migrate all models to creat cooresponding tables
func Migrate() {
	conn := utils.DBConnect()

	// create tables
	conn.AutoMigrate(User{})
	conn.AutoMigrate(Product{})
}
