package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/logger"
	"go-common/app/service/store/config"
	"go-common/library/database/orm"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	c,err := config.GetMySQL()
	if err != nil{
		return
	}

	db = orm.NewMysql(&orm.Config{
		DSN:         c.DSN,
		Active:      c.Active,
		Idle:        c.Idle,
		IdleTimeout: c.IdleTimeout,
	})
	cf = func() {
		logger.Info("close DB")
		db.Close()
	}
	return
}
