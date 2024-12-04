package api

import (
	"context"
	"fmt"
	"log"

	"github.com/roohanyh/lila_p1/proto"
	"github.com/roohanyh/lila_p1/service"
)

const (
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)

type server struct {
	proto.MultiplayerServiceServer
	service service.MultiplayerService
}

func NewServer(service service.MultiplayerService) *server {
	return &server{service: service}
}

func (s *server) GetTopMode(ctx context.Context, req *proto.TopModeRequest) (*proto.TopModeResponse, error) {
	log.Printf("Received GetTopMode request for area_code: %s", req.GetAreaCode())

	if req.GetAreaCode() == "" {
		return nil, fmt.Errorf("area_code is required")
	}

	mode, err := service.GetTopMode(req.GetAreaCode())
	if err != nil {
		return nil, fmt.Errorf("failed to get top mode for area_code %s: %w", req.GetAreaCode(), err)
	}

	resp := &proto.TopModeResponse{
		Mode: mode,
	}
	return resp, nil
}

func (s *server) UpdateSingleMode(ctx context.Context, req *proto.UpdateSingleModeRequest) (*proto.UpdateSingleModeResponse, error) {
	log.Printf("Received UpdateSingleMode request for area_code: %s, mode_name: %s, users: %d",
		req.GetAreaCode(), req.GetModeName(), req.GetUsers())

	if req.GetAreaCode() == "" || req.GetModeName() == "" || req.GetUsers() <= 0 {
		return nil, fmt.Errorf("invalid request parameters")
	}

	err := service.UpdateSingleMode(req.GetAreaCode(), req.GetModeName(), req.GetUsers())
	if err != nil {
		return &proto.UpdateSingleModeResponse{
			Status: StatusFailed,
		}, fmt.Errorf("failed to update mode: %w", err)
	}

	resp := &proto.UpdateSingleModeResponse{
		Status: StatusSuccess,
	}
	return resp, nil
}

func (s *server) RandomizeSingleAreaCode(ctx context.Context, req *proto.RandomizeSingleAreaCodeRequest) (*proto.RandomizeSingleAreaCodeResponse, error) {
	log.Printf("Received RandomizeSingleAreaCode request for area_code: %s", req.GetAreaCode())

	if req.GetAreaCode() == "" {
		return nil, fmt.Errorf("area_code is required")
	}

	err := service.RandomizeSingleAreaCode(req.GetAreaCode(), req.GetSeed())
	if err != nil {
		return &proto.RandomizeSingleAreaCodeResponse{
			Status: StatusFailed,
		}, fmt.Errorf("failed to randomize area code: %w", err)
	}

	resp := &proto.RandomizeSingleAreaCodeResponse{
		Status: StatusSuccess,
	}
	return resp, nil
}

func (s *server) HealthCheck(ctx context.Context, req *proto.EmptyRequest) (*proto.HealthCheckResponse, error) {
	log.Println("Received HealthCheck request")
	return &proto.HealthCheckResponse{Status: "server is up and running..."}, nil
}
