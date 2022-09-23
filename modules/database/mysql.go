package database

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"hometown.com/hometown-serverless-go/modules/common"
)

var (
	_ = godotenv.Load()
	SlaveDatabase = DbConnecttion("SLAVE")
	MasterDatabase = DbConnecttion("MASTER")
)
type DbConfig struct {
	MaxLifetime time.Duration
	MaxIdletime time.Duration
	MaxOpenConn int
	MaxIdleConn int
}

var slaveConfig *DbConfig = &DbConfig{
	MaxLifetime: time.Minute,
	MaxIdletime: time.Minute,
	MaxOpenConn: 60,
	MaxIdleConn: 60,

}

var masterConfig *DbConfig = &DbConfig{
	MaxLifetime: time.Minute,
	MaxIdletime: time.Minute,
	MaxOpenConn: 35,
	MaxIdleConn: 35,
}

func DbConnecttion(env string) *sqlx.DB {
	var config *DbConfig
	switch env {
	case "SLAVE":
		config = slaveConfig
	case "MASTER":
		config = masterConfig
	default:
		config = slaveConfig
	}
	db, dbErr := sqlx.Connect("mysql", os.Getenv("MYSQL_USER_" + env)+":"+os.Getenv("MYSQL_PASSWORD_" + env)+"@tcp("+os.Getenv("MYSQL_HOST_" + env)+")/"+os.Getenv("MYSQL_DATABASE_" + env)+"?charset=utf8mb4&parseTime=True")
	if common.IsError(dbErr) {
		log.Panic(dbErr)
	}
	db.SetConnMaxLifetime(config.MaxLifetime)
	db.SetConnMaxIdleTime(config.MaxIdletime)

	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetMaxOpenConns(config.MaxOpenConn)
	return db
}