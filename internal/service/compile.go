package service

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/pkg/logger"
)

type CompileStep struct {
	log *logger.Logger
	ctx *core.KernelContext
}

var _ core.IStep = (*CompileStep)(nil)

func (s *CompileStep) Name() string {
	return "compile"
}

func (s *CompileStep) Execute() error {
	dir := s.ctx.Directory

	if s.ctx.NumJobs == 0 {
		defaultNumJobs := 1
		s.log.Warn("Number of jobs not configured, defaulting to %d", defaultNumJobs)
		s.ctx.NumJobs = defaultNumJobs
	}

	numJobsStr := fmt.Sprintf("-j%d", s.ctx.NumJobs)

	s.log.Info("Compiling Linux kernel using 'make -C %s %s' ...", dir, numJobsStr)

	cmd := exec.Command("make", "-C", dir, numJobsStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("kernel compilation failed: %w", err)
	}
	return nil
}
