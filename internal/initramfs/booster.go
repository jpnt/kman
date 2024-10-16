package initramfs

import "kman/pkg/logger"

type Booster struct {
	logger 		*logger.Logger
}

func NewBooster(l *logger.Logger) *Booster {
	return &Booster{logger: l}
}

func (b *Booster) Generate() error {
	b.logger.Info("Generating booster initramfs image ...")
	// TODO
	return nil
}
