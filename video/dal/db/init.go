package db

import (
	"github.com/RaymondCode/simple-demo/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&Video{}) {
		return
	}
	if err = m.CreateTable(&Video{}); err != nil {
		panic(err)
	}
}
