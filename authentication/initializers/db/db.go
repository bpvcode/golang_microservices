package db

import (
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	DB       string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Port     int    `env:"DB_PORT"`
	Debug    bool   `env:"DB_DEBUG"`
}

//nolint:gochecknoglobals
var gdb *gorm.DB

//nolint:gochecknoglobals
var doOnce sync.Once

// GetDB return current DB connection.
func GetDB() *gorm.DB {
	doOnce.Do(func() {
		gdb = newDbConnection()
	})

	db, err := gdb.DB()
	if err != nil {
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatal("database ping error")
	}

	return gdb
}

// SetDB will set current DB connection.
func SetDB(db *gorm.DB) {
	gdb = db
}

func (c *config) connectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.User,
		c.Password,
		c.DB,
	)
}

func newDbConnection() *gorm.DB {
	// TODO - CHANGE THIS AND IN HTTP FILE ALSO TO READ AUTOMATIC VARIABLES FROM ENV FILE
	var cfg config = config{
		DB:       os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	var db *gorm.DB
	var err error

	gconfig := &gorm.Config{
		CreateBatchSize:        999,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	}

	db, err = gorm.Open(postgres.Open(cfg.connectionString()), gconfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := db.DB()
	if sqlDB == nil || err != nil {
		log.Fatal("database connection nil")
	}

	if cfg.Debug {
		db = db.Debug()
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatal("database ping error")
	}

	formatter := new(log.TextFormatter)
	log.SetFormatter(formatter)
	formatter.FullTimestamp = true

	log.Info("Successfully connected to database")

	return db

}
