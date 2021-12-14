package Databases

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"web/Config"
)

var (
	db *gorm.DB
)

func initDb() {
	var dbURI string
	var dialector gorm.Dialector
	dbConfig := Config.GetConfig().Db
	dbURI = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
	dialector = mysql.New(mysql.Config{
		DSN:                       dbURI, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := conn.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxIdleTime(time.Second * 600)
	db = conn
}

func init() {
	initDb()
}

func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		initDb()
	}

	return db
}
