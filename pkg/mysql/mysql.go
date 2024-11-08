package mysql

import (
	"github.com/rebirthmonkey/ops/pkg/log"
	"gorm.io/gorm"
)

type DB struct {
	*Config
	DBEngine *gorm.DB
}

var dbInstance *DB

func Init() error {
	opts := NewOptions()
	config := NewConfig()

	if err := opts.ApplyTo(config); err != nil {
		log.Errorln("MySQL Init ApplyTo Config Error: ", err)
		return err
	}

	db, err := config.New()
	if err != nil {
		log.Errorln("MySQL Init Config Error: ", err)
		return err
	}

	dbInstance = db
	return nil
}

func GetUniqueDBInstance() *DB {
	if dbInstance == nil {
		log.Errorln("MySQL GetUniqueDBInstance Error: dbInstance is nil")
		panic("MySQL GetUniqueDBInstance Error: dbInstance is nil")
	}
	return dbInstance
}
