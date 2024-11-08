package mysql

import (
	"fmt"
	"github.com/rebirthmonkey/ops/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func NewConfig() *Config {
	return &Config{
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
		Database: "",
	}
}

func (c *Config) New() (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
	gormInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Logger.Errorln("ConnectMySQL error: ", err, " with dsn: ", dsn)
		panic(err)
	}

	result := gormInstance.Raw("SHOW tables")
	if result.Error != nil {
		log.Errorln("ConnectMySQL error: ", err)
		panic(err)
	}

	db := &DB{
		Config:   c,
		DBEngine: gormInstance,
	}

	return db, nil
}
