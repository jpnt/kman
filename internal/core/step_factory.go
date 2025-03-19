package core

import (
	"github.com/jpnt/kman/pkg/logger"
)

type IStepFactory interface {
	CreateStep(name string, logger logger.ILogger, ctx IKernelContext) (IStep, error)
}
