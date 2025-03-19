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

type PatchStep struct {
	log *logger.Logger
	ctx *core.KernelContext
}

var _ core.IStep = (*PatchStep)(nil)

func (s *PatchStep) Name() string {
	return "patch"
}

func (s *PatchStep) Execute() error {
	s.log.Error("patch step: not implemented yet")

	return nil
}
