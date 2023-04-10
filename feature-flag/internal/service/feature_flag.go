package service

import (
	"context"

	pb "feature-flag/api/admin/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type FeatureFlagService struct {
	pb.UnimplementedFeatureFlagServer

	log log.Logger
}

func NewFeatureFlagService(log log.Logger) *FeatureFlagService {
	return &FeatureFlagService{
		log: log,
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
	log.Infof("s.GetFeatureFlag %+v", req)
	return &pb.GetFeatureFlagReply{}, nil
}

func (s *FeatureFlagService) ListFeatureFlag(ctx context.Context, req *pb.ListFeatureFlagRequest) (*pb.ListFeatureFlagReply, error) {
	return &pb.ListFeatureFlagReply{}, nil
}
