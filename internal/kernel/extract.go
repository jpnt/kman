package kernel

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kman/pkg/utils"
)

func ExtractKernel(kernelPath string) (string, error) {
	kernelDirPath := filepath.Dir(kernelPath)

	// Expected extracted kernel directory
	extractKernelPath := strings.TrimSuffix(kernelPath, filepath.Ext(kernelPath))
	extractKernelPath  = strings.TrimSuffix(extractKernelPath, filepath.Ext(extractKernelPath))

	if _, err := os.Stat(extractKernelPath); err == nil {
		fmt.Printf("Kernel already extracted at: %s\n", extractKernelPath)
		return extractKernelPath, nil
	}

	err := utils.UncompressFile(kernelPath, kernelDirPath)
	if err != nil {
		return "", err
	}
	
	fmt.Printf("Kernel extracted to: %s\n", extractKernelPath)
	return extractKernelPath, nil
}
