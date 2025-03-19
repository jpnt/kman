package core

import (
	"errors"
)

type IKernelContext interface {
	Validate(stepName string) error
}

type KernelContext struct {
	TarballURL    string
	DownloadPath  string
	ArchivePath   string
	Directory     string
	SignatureURL  string
	ConfigOptions []string
	OldConfigPath string
	NumJobs       string
	Initramfs     string
	Bootloader    string
}

var _ IKernelContext = (*KernelContext)(nil)

func NewKernelContext() IKernelContext {
	return &KernelContext{}
}

func (c *KernelContext) Validate(stepName string) error {
	switch stepName {
	case "download":
		if c.TarballURL == "" {
			return errors.New("kernel tarball URL not set")
		}
	case "verify":
		if c.ArchivePath == "" {
			return errors.New("kernel archive path not set")
		}
	case "extract":
		if c.ArchivePath == "" {
			return errors.New("kernel archive path not set")
		}
	}
	return nil
}
