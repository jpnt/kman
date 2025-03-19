package service

import (
	// "errors"
	// "fmt"
	// "os"
	// "os/exec"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/pkg/logger"
	// "github.com/jpnt/kman/pkg/utils"
)

type InstallStep struct {
	log *logger.Logger
	ctx    *core.KernelContext
}

var _ core.IStep = (*InstallStep)(nil)

func (s *InstallStep) Name() string {
	return "install"
}

func (s *InstallStep) Execute() error {
	s.log.Warn("install step: not implemented yet")

	return nil
}
