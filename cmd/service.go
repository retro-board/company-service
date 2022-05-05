package main

import (
	"fmt"
	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/retro-board/company-service/internal/config"
	"github.com/retro-board/company-service/internal/service"
)

var (
	BuildVersion = "dev"
	BuildHash    = "unknown"
)

func main() {
	bugLog.Local().Info("Starting Service")
	bugLog.Local().Info(fmt.Sprintf("Version: %s, Hash: %s", BuildVersion, BuildHash))

	cfg, err := config.Build()
	if err != nil {
		_ = bugLog.Errorf("config: %v", err)
		return
	}

	s := &service.Service{
		Config: cfg,
	}

	if err := s.Start(); err != nil {
		_ = bugLog.Errorf("start service: %v", err)
		return
	}
}
