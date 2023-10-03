package bootstrap

import (
	"database/sql"
	"fmt"
	"gin-clean-arch/domain"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(env *Env) *gorm.DB {
	dbHost := env.DbHost
	dbPort := env.DbPort
	dbUser := env.DbUser
	dbName := env.DbName
	dbPass := env.DbPass

	dsn := fmt.Sprintf("host=%v user=%v dbname=%v port=%v password=%v sslmode=disable",
		dbHost, dbUser, dbName, dbPort, dbPass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("can't connect %v", err.Error())
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	return db
}

func GenerateDatabase(db *gorm.DB) {
	var sqlTest *sql.DB
	var err error

	sqlTest, err = db.DB()

	if err != nil {
		fmt.Printf("postgres problem %v", err.Error())
	}

	err = sqlTest.Ping()

	if err != nil {
		fmt.Printf("not connected - %v", err.Error())
		return
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		fmt.Printf("error in migration process")
	}
}
