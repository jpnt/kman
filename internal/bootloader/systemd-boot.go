package bootloader

import "kman/pkg/logger"

type SystemdBoot struct {
	logger *logger.Logger
}

var _ IBootloader = (*SystemdBoot)(nil)

func NewSystemdBoot(l *logger.Logger) *SystemdBoot {
	return &SystemdBoot{logger: l}
}

func (s *SystemdBoot) Configure() error {
	s.logger.Info("Configuring systemd-boot bootloader ...")
	// TODO
	return nil
}
