package initramfs

import "kman/pkg/logger"

type Dracut struct {
	logger 		*logger.Logger
}

func NewDracut(l *logger.Logger) *Dracut {
	return &Dracut{logger: l}
}

func (d *Dracut) Generate() error {
	d.logger.Info("Generating mkinicpio initramfs image ...")
	// TODO
	return nil
}
