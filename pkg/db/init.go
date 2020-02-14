package db

import (
	"rest-shell/pkg/utils/syslog"
	"fmt"
	"os"
	"time"

	"rest-shell/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DbConn *gorm.DB

func InitDB() (*gorm.DB, error) {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("username"), os.Getenv("database"), os.Getenv("password"))
	LOG.Info("dburl:", dbUrl)
	DbConn, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	DbConn.DB().SetMaxIdleConns(10)
	DbConn.DB().SetMaxOpenConns(100)
	DbConn.DB().SetConnMaxLifetime(time.Hour)

	DbConn.AutoMigrate(&model.Rolebinding{})
	return DbConn, nil
}
