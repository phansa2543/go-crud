package databases

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/phansa2543/go-crud/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB() Database {
	conf := config.LoadConfig()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
		  SlowThreshold:              time.Second,
		  LogLevel:                   logger.Info,
		  IgnoreRecordNotFoundError: true,
		  ParameterizedQueries:      true,
		  Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			conf.PostgresDB.Host,
			conf.PostgresDB.Port,
			conf.PostgresDB.User,
			conf.PostgresDB.Password,
			conf.PostgresDB.DBName,
			conf.PostgresDB.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	return &PostgresDB{db: db}
}

func (p *PostgresDB) ConnectDB() *gorm.DB {
	return p.db
}