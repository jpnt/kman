package initramfs

import "kman/pkg/logger"

type Mkinitcpio struct {
	logger *logger.Logger
}

var _ IInitramfs = (*Mkinitcpio)(nil)

func NewMkinitcpio(l *logger.Logger) *Mkinitcpio {
	return &Mkinitcpio{logger: l}
}

func (m *Mkinitcpio) Generate() error {
	m.logger.Info("Generating mkinicpio initramfs image ...")
	// TODO
	return nil
}
