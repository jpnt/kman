package kernel

import (
	"github.com/jpnt/kman/pkg/progress"
	"github.com/jpnt/kman/pkg/utils"
)

func DownloadKernel(sourceURL, destPath string) (string, error) {
	pb := &progress.WriteCounter{}

	kernelPath, err := utils.DownloadFile(sourceURL, destPath, pb)
	if err != nil {
		return "", err
	}

	return kernelPath, nil
}
