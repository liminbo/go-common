package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/logger"
	"strings"
	"time"
)

// Config mysql config.
type Config struct {
	DSN          string          // write data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  int   // connect max life time.
}

type Gorm struct {

}

func (g Gorm) Print(v ...interface{}) {
	logger.Infof(strings.Repeat("%v ", len(v)), v...)
}

func NewMysql(c *Config) (db *gorm.DB){
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		logger.Errorf("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	db.SingularTable(true)
	db.SetLogger(Gorm{})

	return
}