package kernel

import (
	// "errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/jpnt/kman/pkg/logger"
	// "github.com/jpnt/kman/pkg/utils"
)

type CompileStep struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ IStep = (*CompileStep)(nil)

func (s *CompileStep) String() string {
	return "[CompileStep]"
}

func (s *CompileStep) Execute() error {
	dir := s.ctx.directory
	// njobs := s.ctx.njobs // TODO
	njobs := 1
	njobsStr := fmt.Sprintf("-j%d", njobs)

	s.logger.Info("Compiling Linux kernel: 'make -C %s %s' ...", dir, njobsStr)
	
	cmd := exec.Command("make", "-C", dir, njobsStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("kernel compilation failed: %w", err)
	}

	s.logger.Info("Compiled Linux kernel")

	return nil
}
