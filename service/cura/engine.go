package cura

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
)

type Engine struct {
	binary   string
	basepath string
}

func NewEngine(binary string, basepath string) *Engine {
	if len(strings.Trim(binary, " ")) == 0 {
		return nil
	}

	err := exec.Command(binary, "help").Run()
	if err != nil {
		slog.Warn("The binary path of CuraEngine cannot be executed")
		return nil
	}

	if len(strings.Trim(basepath, " ")) != 0 {
		slog.Info(fmt.Sprintf("Load CuraEngine configuration from %s", basepath))
	}

	return &Engine{
		binary:   binary,
		basepath: basepath,
	}
}
