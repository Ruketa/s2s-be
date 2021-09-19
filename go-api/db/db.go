package db

import (
	"goapi/entity"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/ini.v1"
)

var (
	db  *gorm.DB
	err error
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

	db, err = gorm.Open("postgres",
		"host="+config.Host+
			" port="+strconv.Itoa(config.Port)+
			" user="+config.User+
			" dbname="+config.Database+
			" password="+config.Password+
			" sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	AutoMigrate()
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}

// CloseDB closes the database connection
func CloseDB() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// AutoMigrate will attempt to automatically migrate all tables
func AutoMigrate() {
	db.AutoMigrate(&entity.Questionnaire{}, &entity.PresentationPlan{}, &entity.StudySession{})
}
