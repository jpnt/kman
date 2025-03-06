package kernel

import (
	// "errors"
	// "fmt"
	// "os"
	// "os/exec"

	"github.com/jpnt/kman/pkg/logger"
	// "github.com/jpnt/kman/pkg/utils"
)

type PatchStep struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ IStep = (*PatchStep)(nil)

func (s *PatchStep) String() string {
	return "[PatchStep]"
}

func (s *PatchStep) Execute() error {
	s.logger.Error("patch step: not implemented yet")
	
	return nil
}
