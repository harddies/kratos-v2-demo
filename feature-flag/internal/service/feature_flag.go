package service

import (
	"context"
	"feature-flag/internal/biz"

	pb "feature-flag/api/admin/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type FeatureFlagService struct {
	pb.UnimplementedFeatureFlagServer

	log log.Logger

	featureFlagDS *biz.FeatureFlagDS
}

func NewFeatureFlagService(log log.Logger, featureFlagDS *biz.FeatureFlagDS) *FeatureFlagService {
	return &FeatureFlagService{
		log:           log,
		featureFlagDS: featureFlagDS,
	}
}

func (s *FeatureFlagService) CreateFeatureFlag(ctx context.Context, req *pb.CreateFeatureFlagRequest) (*pb.CreateFeatureFlagReply, error) {
	return &pb.CreateFeatureFlagReply{}, nil
}

func (s *FeatureFlagService) UpdateFeatureFlag(ctx context.Context, req *pb.UpdateFeatureFlagRequest) (*pb.UpdateFeatureFlagReply, error) {
	return &pb.UpdateFeatureFlagReply{}, nil
}

func (s *FeatureFlagService) DeleteFeatureFlag(ctx context.Context, req *pb.DeleteFeatureFlagRequest) (*pb.DeleteFeatureFlagReply, error) {
	return &pb.DeleteFeatureFlagReply{}, nil
}

func (s *FeatureFlagService) GetFeatureFlag(ctx context.Context, req *pb.GetFeatureFlagRequest) (*pb.GetFeatureFlagReply, error) {
	featureFlagDO, err := s.featureFlagDS.GetFeatureFlag(ctx, req.GetFeatureFlagId())
	if err != nil {
		return nil, err
	}
	log.Infof("s.GetFeatureFlag req(%+v) featureFlagDO(%+v)", req, featureFlagDO)
	return &pb.GetFeatureFlagReply{}, nil
}

func (s *FeatureFlagService) ListFeatureFlag(ctx context.Context, req *pb.ListFeatureFlagRequest) (*pb.ListFeatureFlagReply, error) {
	return &pb.ListFeatureFlagReply{}, nil
}
