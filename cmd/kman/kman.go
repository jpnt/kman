package main

import (
	"fmt"
	"os"

	"kman/internal/kernel"
	"kman/pkg/logger"
	"kman/pkg/progress"
	"kman/pkg/utils"
)

const (
	KERNEL_SRC = "/home/jpnt/Coding/Linux/kman-kernels"
)

func main() {
	log := logger.NewLogger()
	defer log.Info("Exited")

	log.Info("Fetching kernel versions...")
	kernels, err := kernel.ListKernels()
	if err != nil {
		log.Error("Failed to fetch kernels: %s", err.Error())
		os.Exit(1)
	}

	// Prompt to select kernel
	n_kernels := len(kernels) - 1
	var i int
	for {
		fmt.Printf("Please select a kernel to download (0-%d): ", n_kernels)
		fmt.Scanf("%d", &i)
		if i >= 0 && i <= n_kernels {
			break
		}
	}

	selectedKernel := kernels[i]
	log.Info("Selected kernel: %s", selectedKernel.Title)

	if selectedKernel.SourceTarball == "" {
		log.Error("This kernel version does not have a source tarball")
		return
	}

	if selectedKernel.PGPSignature == "" {
		log.Warn("This kernel version does not have a PGP signature")
		confirm := utils.ConfirmAction("Are you okay with this? (y/N)")
		if !confirm {
			return
		}
	}

	kernelPath, err := kernel.DownloadKernel(selectedKernel.SourceTarball, KERNEL_SRC)
	if err != nil {
		log.Error("Failed to download kernel: %s", err.Error())
	}

	if err = kernel.VerifyKernelSignature(selectedKernel.PGPSignature, kernelPath); err != nil {
		log.Error("Failed to verify kernel: %s", err)
		confirm := utils.ConfirmAction("Are you okay with this? (y/N)")
		if !confirm {
			return
		}
	}

	_, err = kernel.UncompressKernel(kernelPath)
	if err != nil {
		log.Error("Failed to uncompress kernel: %s", err.Error())
	}
}
