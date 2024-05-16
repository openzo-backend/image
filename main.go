package main

import (
	"fmt"
	"log"

	"github.com/tanush-128/openzo_backend/image/config"
	"github.com/tanush-128/openzo_backend/image/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to load config: %w", err))
	}

	service.GrpcServer(
		cfg,
		&service.Server{},
	)

}
