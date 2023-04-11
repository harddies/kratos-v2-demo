package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type IFeatureFlagRepo interface {
	GetByFeatureFlagId(c context.Context, featureFlagId string) (featureFlagDo *FeatureFlagDO, err error)
}

type FeatureFlagDO struct {
	FeatureFlagId string
	FeatureKey    string
}

func NewFeatureFlagDS(featureFlagRepo IFeatureFlagRepo, log log.Logger) *FeatureFlagDS {
	return &FeatureFlagDS{
		featureFlagRepo: featureFlagRepo,
		log:             log,
	}
}

type FeatureFlagDS struct {
	featureFlagRepo IFeatureFlagRepo
	log             log.Logger
}

func (s *FeatureFlagDS) GetFeatureFlag(c context.Context, featureFlagId string) (featureFlagDo *FeatureFlagDO, err error) {
	if featureFlagDo, err = s.featureFlagRepo.GetByFeatureFlagId(c, featureFlagId); err != nil {
		return
	}
	return
}
