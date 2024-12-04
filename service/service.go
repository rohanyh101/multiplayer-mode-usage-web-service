package service

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/roohanyh/lila_p1/cache"
	"github.com/roohanyh/lila_p1/database"
	"github.com/roohanyh/lila_p1/models"
	"github.com/roohanyh/lila_p1/proto"
)

type MultiplayerService interface {
	GetTopMode(areaCode string) (string, error)
	UpdateSingleMode(areaCode, modeName string, users int32) error
	RandomizeSingleAreaCode(areaCode string, seed int64) error
}

func GetTopMode(areaCode string) (*proto.Mode, error) {
	if areaCode == "" {
		return nil, fmt.Errorf("areaCode is required")
	}

	mode, err := cache.GetTopMode(areaCode)
	if err == nil && mode != nil {
		return mode, nil
	}

	mode, err = database.GetTopMode(areaCode)
	if err != nil {
		return nil, fmt.Errorf("failed to get top mode for areaCode %s: %w", areaCode, err)
	}

	cacheMode := &models.CacheMode{
		Name:  mode.Name,
		Users: mode.Users,
	}

	if err := cache.SetTopMode(areaCode, *cacheMode); err != nil {
		log.Printf("Failed to update cache for areaCode %s: %v", areaCode, err)
	}

	return mode, nil
}

func UpdateSingleMode(areaCode, modeName string, userCount int32) error {
	if areaCode == "" || modeName == "" || userCount <= 0 {
		return fmt.Errorf("invalid request parameters")
	}

	mongoMode, err := database.GetModeByName(modeName)
	if err != nil {
		return fmt.Errorf("failed to fetch mode %s: %w", modeName, err)
	}

	if err := database.UpdateSingleMode(areaCode, mongoMode.ID, userCount); err != nil {
		return fmt.Errorf("failed to update mode %s for areaCode %s: %w", modeName, areaCode, err)
	}

	topMode, err := database.GetTopMode(areaCode)
	if err != nil {
		return fmt.Errorf("failed to fetch top mode for areaCode %s: %w", areaCode, err)
	}

	cacheTopMode := &models.CacheMode{
		Name:  topMode.Name,
		Users: topMode.Users,
	}

	if err := cache.SetTopMode(areaCode, *cacheTopMode); err != nil {
		log.Printf("Failed to update cache for areaCode %s: %v", areaCode, err)
	}

	return nil
}

func RandomizeSingleAreaCode(areaCode string, seed int32) error {
	if areaCode == "" {
		return fmt.Errorf("areaCode is required")
	}

	areaCodeData, err := database.GetAreaCode(areaCode)
	if err != nil {
		return fmt.Errorf("failed to fetch area code %s: %w", areaCode, err)
	}

	if areaCodeData.AreaCode == "" {
		return fmt.Errorf("area code %s not found", areaCode)
	}

	for i := range areaCodeData.ModeTraffic {
		areaCodeData.ModeTraffic[i].Users = int32(rand.Intn(int(seed)))
	}

	if err := database.UpdateModeTraffic(areaCodeData); err != nil {
		return fmt.Errorf("failed to update mode traffic for areaCode %s: %w", areaCode, err)
	}

	topMode, err := database.GetTopMode(areaCode)
	if err != nil {
		return fmt.Errorf("failed to fetch top mode for areaCode %s: %w", areaCode, err)
	}

	cacheTopMode := &models.CacheMode{
		Name:  topMode.Name,
		Users: topMode.Users,
	}
	if err := cache.SetTopMode(areaCode, *cacheTopMode); err != nil {
		log.Printf("Failed to update cache for areaCode %s: %v", areaCode, err)
	}

	return nil
}
