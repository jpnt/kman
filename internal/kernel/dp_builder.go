package kernel

import (
	"fmt"

	"github.com/jpnt/kman/pkg/logger"
)

type IKernelBuilder interface {
	WithList() IKernelBuilder
	WithDownload() IKernelBuilder
	WithVerify() IKernelBuilder
	WithExtract() IKernelBuilder
	WithConfigure() IKernelBuilder
	// WithPatch() IKernelBuilder // not implemented yet
	// WithCompile() IKernelBuilder // not implemented yet
	// WithInstall() IKernelBuilder // not implemented yet
	WithDefault() IKernelBuilder
	Build() (IKernelFacade, error)
}

type KernelBuilder struct {
	logger *logger.Logger
	cm     *CommandManager
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ IKernelBuilder = (*KernelBuilder)(nil)

func NewKernelBuilder(l *logger.Logger) IKernelBuilder {
	sharedCtx := &KernelContext{}
	return &KernelBuilder{
		logger: l,
		cm:     NewCommandManager(),
		ctx:    sharedCtx,
	}
}

func (kb *KernelBuilder) addCommand(cmd ICommand) IKernelBuilder {
	kb.cm.AddCommand(cmd)
	return kb
}

// Still a bit clunky... Maybe look at Factory pattern or registry to map string to command?
func (kb *KernelBuilder) WithList() IKernelBuilder {
	return kb.addCommand(&ListCommand{logger: kb.logger, ctx: kb.ctx})
}

func (kb *KernelBuilder) WithDownload() IKernelBuilder {
	return kb.addCommand(&DownloadCommand{logger: kb.logger, ctx: kb.ctx})
}

func (kb *KernelBuilder) WithVerify() IKernelBuilder {
	return kb.addCommand(&VerifyCommand{logger: kb.logger, ctx: kb.ctx})
}

func (kb *KernelBuilder) WithExtract() IKernelBuilder {
	return kb.addCommand(&ExtractCommand{logger: kb.logger, ctx: kb.ctx})
}

func (kb *KernelBuilder) WithConfigure() IKernelBuilder {
	return kb.addCommand(&ConfigureCommand{logger: kb.logger, ctx: kb.ctx})
}

func (kb *KernelBuilder) WithDefault() IKernelBuilder {
	kb.
		WithList().
		WithDownload().
		WithVerify().
		WithExtract().
		WithConfigure()
	return kb
}

func (kb *KernelBuilder) Build() (IKernelFacade, error) {
	if len(kb.cm.commands) == 0 {
		return nil, fmt.Errorf("no commands were configured in the KernelBuilder")
	}
	return NewKernelFacade(kb.logger, kb.cm, kb.ctx), nil
}
