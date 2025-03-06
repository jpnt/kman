package kernel

import (
	// "errors"
	// "fmt"
	// "os"
	// "os/exec"

	"github.com/jpnt/kman/pkg/logger"
	// "github.com/jpnt/kman/pkg/utils"
)

type InstallStep struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ IStep = (*InstallStep)(nil)

func (s *InstallStep) String() string {
	return "[InstallStep]"
}

func (s *InstallStep) Execute() error {
	s.logger.Warn("install step: not implemented yet")
	
	return nil
}
