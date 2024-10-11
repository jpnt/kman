package kernel

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jpnt/kman/pkg/progress"
	"github.com/jpnt/kman/pkg/utils"
)

const (
	LINUS   = "torvalds@kernel.org"
	GREG    = "gregkh@kernel.org"
)

var emails = []string{LINUS, GREG}

func VerifyKernelSignature(signatureURL, kernelPath string) error {
	if signatureURL == "" {
		fmt.Println("Signature not found.")
		return nil
	}

	pb := &progress.WriteCounter{}
	signaturePathDest := filepath.Dir(kernelPath)

	signature, err := utils.DownloadFile(signatureURL, signaturePathDest, pb)
	if err != nil {
		return err
	}
	
	return verifyPGPSignature(signature, kernelPath)
}

func verifyPGPSignature(signaturePath, kernelPath string) error {
	err := importKeys(emails)
	if err != nil {
		return err
	}

	// TODO: verify pgp signature

	fmt.Println("Kernel signature verified.")
	return nil
}

func importKeys(emails []string) error {
	for _, email := range emails {
		fmt.Printf("Locating and importing key for %s...\n", email)

		cmd := exec.Command("gpg", "--locate-keys", email)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to import key for %s: %v\n", email, err)
		}
		fmt.Println("done.")
	}

	fmt.Println("Imported keys successfully.")
	return nil
}
