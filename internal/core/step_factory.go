package core

import (
	"github.com/jpnt/kman/pkg/logger"
)

type IStepFactory interface {
	CreateStep(name string, log logger.ILogger, ctx IKernelContext) (IStep, error)
}
