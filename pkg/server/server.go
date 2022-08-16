package server

import (
	"context"

	"github.com/jwt-auth/pkg/config"
	"github.com/jwt-auth/pkg/services"
)

func Execute(ctx context.Context) error {
	config, err := config.LoadConfig("../")

	if err != nil {
		return err
	}

	return services.InitServices(config)
}
