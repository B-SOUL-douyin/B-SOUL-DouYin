package db

import (
	"github.com/B-SOUL-douyin/B-SOUL-DouYin/pkg/constants"
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
	if m.HasTable(&VideoModel{}) {
		return
	}
	if err = m.CreateTable(&VideoModel{}); err != nil {
		panic(err)
	}
}
