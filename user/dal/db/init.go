package db

import (
	"github.com/B-SOUL-douyin/B-SOUL-DouYin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	var user = constants.User{}
	if m.HasTable(&user) {
		return
	}
	if err = m.CreateTable(&user); err != nil {
		panic(err)
	}
}
