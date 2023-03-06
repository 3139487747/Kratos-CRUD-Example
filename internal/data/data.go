package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shop/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewCrudRepo)

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	db, err := gorm.Open(mysql.Open(c.GetDatabase().GetSource()), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}
