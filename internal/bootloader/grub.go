package bootloader

import "kman/pkg/logger"

type Grub struct {
	logger *logger.Logger
}

var _ IBootloader = (*Grub)(nil)

func NewGRUB(l *logger.Logger) *Grub {
	return &Grub{logger: l}
}

func (g *Grub) Configure() error {
	g.logger.Info("Configuring GRUB bootloader ...")
	// TODO
	return nil
}
