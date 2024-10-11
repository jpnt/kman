package bootloader

import (
	"github.com/jpnt/kman/pkg/logger"
)

type GRUB struct {
	logger *logger.Logger
}

func NewGRUB(l *logger.Logger) *GRUB {
	return &GRUB{logger: l}
}

func (g *GRUB) Configure() error {
	g.logger.Info("Configuring GRUB bootloader...")
	// TODO
	return nil
}
