package service

import (
	// "errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/pkg/logger"
	// "github.com/jpnt/kman/pkg/utils"
)

type CompileStep struct {
	log *logger.Logger
	ctx    *core.KernelContext
}

var _ core.IStep = (*CompileStep)(nil)

func (s *CompileStep) Name() string {
	return "compile"
}

// TODO
func (s *CompileStep) Execute() error {
	dir := s.ctx.Directory
	// njobs := s.ctx.njobs // TODO
	njobs := 1
	njobsStr := fmt.Sprintf("-j%d", njobs)

	s.log.Info("Compiling Linux kernel: 'make -C %s %s' ...", dir, njobsStr)

	cmd := exec.Command("make", "-C", dir, njobsStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("kernel compilation failed: %w", err)
	}

	s.log.Info("Compiled Linux kernel")

	return nil
}
