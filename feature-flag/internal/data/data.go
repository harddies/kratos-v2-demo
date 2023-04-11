package data

import (
	"feature-flag/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewFeatureFlagRepoImpl)

// Data .
type Data struct {
	orm *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (d *Data, cleanup func(), err error) {
	cleanup = func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	d = &Data{}
	if d.orm, err = gorm.Open(c.Database.Driver, c.Database.Source); err != nil {
		return
	}
	d.orm.SingularTable(true)
	d.orm.LogMode(true)
	return
}
