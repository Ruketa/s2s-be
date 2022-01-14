package db

import (
	model "goapi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gopkg.in/ini.v1"
)

var (
	db *gorm.DB
)

type ConfigList struct {
	Host     string
	Database string
	User     string
	Port     int
	Password string
}

// Load config file
func LoadConfig() *ConfigList {
	cfg, err := ini.Load("./db/dbconfig.ini")
	if err != nil {
		panic(err)
	}
	return &ConfigList{
		Host:     cfg.Section("db").Key("host").String(),
		Database: cfg.Section("db").Key("database").String(),
		User:     cfg.Section("db").Key("user").String(),
		Port:     cfg.Section("db").Key("port").MustInt(),
		Password: cfg.Section("db").Key("password").String(),
	}
}

// InitDB initializes the database connection
func InitDB() {

	config := LoadConfig()

	username := config.User
	password := config.Password
	dbName := config.Database
	dbHost := config.Host
	dns := "user=" + username + " password=" + password + " host=" + dbHost + " dbname=" + dbName + " sslmode=disable"
	gormDB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = gormDB

	AutoMigrate()
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}

// CloseDB closes the database connection
func CloseDB() {
	database, err := db.DB()
	if err != nil {
		panic(err)
	}
	database.Close()
}

// AutoMigrate will attempt to automatically migrate all tables
func AutoMigrate() {
	db.AutoMigrate(&model.Questionnaire{}, &model.PresentationPlan{}, &model.StudySession{})
}
