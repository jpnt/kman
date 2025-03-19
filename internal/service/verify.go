package service

// TODO: looks super ugly (like list.go) try to embed this into the binary?

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/pkg/logger"
	"github.com/jpnt/kman/pkg/progress"
	"github.com/jpnt/kman/pkg/utils"
)

type VerifyStep struct {
	log *logger.Logger
	ctx *core.KernelContext
}

var _ core.IStep = (*VerifyStep)(nil)

func (s *VerifyStep) Name() string {
	return "verify"
}

func (s *VerifyStep) Execute() error {
	if s.ctx.SignatureURL == "" {
		s.log.Warn("Skipping PGP signature verification step ...")
		return nil
	}

	pb := &progress.WriteCounter{}
	signaturePathDest := filepath.Dir(s.ctx.DownloadPath)

	s.log.Info("Downloading signature from %s ...", s.ctx.SignatureURL)

	signaturePath, err := utils.DownloadFile(s.ctx.SignatureURL, signaturePathDest, pb)
	if err != nil {
		return fmt.Errorf("failed to download signature: %w", err)
	}
	s.log.Info("Downloaded signature to %s", signaturePath)

	err = s.verifyKernelPGPSignature(signaturePath, s.ctx.ArchivePath)
	if err != nil {
		return err
	}

	if err := utils.RemoveFile(signaturePath); err != nil {
		return fmt.Errorf("failed to remove signature file: %w", err)
	}
	s.log.Info("Removed signature file: %s", signaturePath)

	return nil
}

var emails = []string{"torvalds@kernel.org", "gregkh@kernel.org"}

// https://www.kernel.org/signature.html
func (s *VerifyStep) verifyKernelPGPSignature(signaturePath, kernelPath string) error {
	s.log.Info("Verifying Linux kernel signature ...")

	err := s.importKeys(emails)
	if err != nil {
		return err
	}

	if filepath.Ext(kernelPath) != ".xz" {
		return fmt.Errorf("the kernel archive is not *.xz")
	}

	unxzCmd := exec.Command("xz", "-cd", kernelPath)
	gpgCmd := exec.Command("gpg", "--verify", signaturePath, "-")

	pipe, err := unxzCmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create pipe: %w", err)
	}
	gpgCmd.Stdin = pipe
	gpgCmd.Stdout = os.Stdout
	gpgCmd.Stderr = os.Stderr

	if err := unxzCmd.Start(); err != nil {
		return fmt.Errorf("failed to start decompression: %w", err)
	}

	if err := gpgCmd.Run(); err != nil {
		return fmt.Errorf("signature verification failed: %w", err)
	}

	if err := unxzCmd.Wait(); err != nil {
		return fmt.Errorf("decompression failed: %v", err)
	}

	s.log.Info("Linux kernel signature verified")
	return nil
}

func (s *VerifyStep) importKeys(emails []string) error {
	for _, email := range emails {
		s.log.Info("Locating and importing key for %s ...", email)

		cmd := exec.Command("gpg", "--locate-keys", email)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to import key for %s: %v\n", email, err)
		}
		s.log.Info("Imported key for %s", email)
	}

	return nil
}
