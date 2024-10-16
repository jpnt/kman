package kernel

import (
	"kman/pkg/progress"
	"kman/pkg/utils"
)

func DownloadKernel(sourceURL, destPath string) (string, error) {
	p := &progress.WriteCounter{}

	kernelPath, err := utils.DownloadFile(sourceURL, destPath, p)
	if err != nil {
		return "", err
	}

	return kernelPath, nil
}
