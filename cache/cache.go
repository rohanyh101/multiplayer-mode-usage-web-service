package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/roohanyh/lila_p1/models"
	"github.com/roohanyh/lila_p1/proto"
)

func GetTopMode(ac string) (*proto.Mode, error) {
	if dbClient == nil {
		return nil, fmt.Errorf("cache client is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	modeBytes, err := dbClient.Get(ctx, ac).Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cache for area code %s: %w", ac, err)
	}

	var mode models.CacheMode
	if err := json.Unmarshal(modeBytes, &mode); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cache data for area code %s: %w", ac, err)
	}

	return &proto.Mode{
		Name:  mode.Name,
		Users: mode.Users,
	}, nil
}

func SetTopMode(ac string, mode models.CacheMode) error {
	if dbClient == nil {
		return fmt.Errorf("cache client is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	modeBytes, err := json.Marshal(mode)
	if err != nil {
		return fmt.Errorf("failed to marshal mode for area code %s: %w", ac, err)
	}

	const expiry = 5 * time.Minute
	if err := dbClient.Set(ctx, ac, modeBytes, expiry).Err(); err != nil {
		return fmt.Errorf("failed to set cache for area code %s: %w", ac, err)
	}

	return nil
}
