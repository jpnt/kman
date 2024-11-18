package kernel

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"kman/pkg/logger"
	"kman/pkg/progress"
	"kman/pkg/utils"
)

type VerifyCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ ICommand = (*VerifyCommand)(nil)

func (c *VerifyCommand) Execute() error {
	if c.ctx.signatureURL == "" {
		c.logger.Warn("Skipping verify PGP signature command")
		return nil
	}

	pb := &progress.WriteCounter{}
	signaturePathDest := filepath.Dir(c.ctx.downloadPath)
	c.logger.Info("Downloading signature from %s", c.ctx.signatureURL)
	signaturePath, err := utils.DownloadFile(c.ctx.signatureURL, signaturePathDest, pb)
	if err != nil {
		return fmt.Errorf("failed to download signature: %w", err)
	}

	c.logger.Info("Downloaded signature to %s", signaturePath)

	err = verifyKernelPGPSignature(signaturePath, c.ctx.archivePath)
	if err != nil {
		return err
	}

	c.logger.Info("Linux kernel signature verified")
	return nil
}

var emails = []string{"torvalds@kernel.org", "gregkh@kernel.org"}

// https://www.kernel.org/signature.html
func verifyKernelPGPSignature(signaturePath, kernelPath string) error {
	err := importKeys(emails)
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

	return nil
}

func importKeys(emails []string) error {
	for _, email := range emails {
		fmt.Printf("Locating and importing key for %s...\n", email)

		cmd := exec.Command("gpg", "--locate-keys", email)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to import key for %s: %v\n", email, err)
		}

		fmt.Printf("Key for %s imported successfully.\n", email)
	}

	return nil
}
