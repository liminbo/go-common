package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/logger"
)


//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
}

// dao dao.
type dao struct {
	db          *gorm.DB
}

// New new a dao and return.
func New(db *gorm.DB) (d Dao, cf func(), err error) {
	return newDao(db)
}

func newDao(db *gorm.DB) (d *dao, cf func(), err error) {
	d = &dao{
		db: db,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	logger.Info("close dao")
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
