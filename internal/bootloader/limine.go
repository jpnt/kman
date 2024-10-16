package bootloader

import "kman/pkg/logger"

type Limine struct {
	logger *logger.Logger
}

func NewLimine(l *logger.Logger) *Limine {
	return &Limine{logger: l}
}

func (g *Limine) Configure() error {
	g.logger.Info("Configuring Limine bootloader ...")
	// TODO
	return nil
}
