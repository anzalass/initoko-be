package database

import (
	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func InitDatabase() *gorm.DB {
// 	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.InitConfig().DBUser, config.InitConfig().DBPass, config.InitConfig().DBHost, config.InitConfig().DBPort, config.InitConfig().DBName)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}

// 	logrus.Error("Database : Connect to MySQL Successfully")
// 	return db
// }

func InitDatabase() *gorm.DB {
	var dsn = "user=postgres.tvxidjlvwoqmuajidtsw password=Baseball-27-08-03~ host=aws-0-ap-southeast-1.pooler.supabase.com port=5432 dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Error("Database : Connect to MySQL Successfully")
	return db
}
