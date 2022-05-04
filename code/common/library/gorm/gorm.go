package gorm

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	orm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *orm.DB

func New(conf string) {
	db, err := orm.Open(mysql.Open(conf), getGormConfig())
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	_db = db
	// err = db.AutoMigrate(&orm.Model{}) // auto generate table ddl
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }
}

func DB() *orm.DB {
	return _db
}

func getGormConfig() *orm.Config {
	// slowLogger := logger.New(
	// 	// standard output
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,   // slow SQL threshold
	// 		LogLevel:      logger.Silent, // log level
	// 		Colorful:      false,         // disable color print
	// 	},
	// )
	return &orm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// Logger: slowLogger,
	}
}
