package data

import (
	"context"
	"github.com/jinzhu/gorm"

	"feature-flag/internal/biz"
)

type FeatureFlag struct {
	ID            int64  `gorm:"id"`
	FeatureFlagId string `gorm:"feature_flag_id"`
	FeatureKey    string `gorm:"feature_key"`
}

var _ biz.IFeatureFlagRepo = (*FeatureFlagRepoImpl)(nil)

func NewFeatureFlagRepoImpl(d *Data) biz.IFeatureFlagRepo {
	return &FeatureFlagRepoImpl{
		d: d,
	}
}

type FeatureFlagRepoImpl struct {
	d *Data
}

func (r *FeatureFlagRepoImpl) GetByFeatureFlagId(c context.Context, featureFlagId string) (featureFlagDo *biz.FeatureFlagDO, err error) {
	po := &FeatureFlag{}
	if err = r.d.orm.Where("feature_flag_id = ?", featureFlagId).First(po).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}

	featureFlagDo = &biz.FeatureFlagDO{
		FeatureFlagId: po.FeatureFlagId,
		FeatureKey:    po.FeatureKey,
	}
	return
}
