package main

import (
	"fmt"

	"github.com/kardianos/service"
)

func main() {
	svcConfig := &service.Config{
		Name:        "WinSenseConnect",
		DisplayName: "MQTT Powershell Automation Service",
		Description: "Listens for MQTT messages and runs PowerShell scripts",
	}

	prg, err := newProgram()
	if err != nil {
		fmt.Printf("Failed to create program: %v\n", err)
		return
	}
	defer prg.logger.Close()

	s, err := service.New(prg, svcConfig)
	if err != nil {
		prg.logger.Error(fmt.Sprintf("Failed to create service: %v", err))
		return
	}

	prg.logger.Debug("Service created, running service...")

	err = s.Run()
	if err != nil {
		prg.logger.Error(fmt.Sprintf("Service failed: %v", err))
		return
	}

	prg.logger.Debug("Service run completed")
}
