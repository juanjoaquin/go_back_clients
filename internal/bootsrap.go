package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/juanjoaquin/go-domain-clients/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		err := db.AutoMigrate(
			&domain.User{},
			&domain.Client{},
			&domain.Category{},
			&domain.Product{},
			&domain.Payment{},
			&domain.PaymentMethod{},
			&domain.CheckingAccount{},
			&domain.MovementCheckingAccount{},
			&domain.ProductStockMovement{},
			&domain.Sale{},
			&domain.SaleDetail{},
		)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func InitLogger() *log.Logger {
	return log.New(os.Stdout, "[DB] ", log.LstdFlags|log.Lshortfile)
}
