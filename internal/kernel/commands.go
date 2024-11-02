package kernel

import (
	// "kman/pkg/progress"
	"kman/pkg/utils"
)

type DownloadCommand struct {}

func (c *DownloadCommand) Execute() error {
	if (!utils.ConfirmAction("Do you wish to download a kernel? (y/N)")) {
		return nil
	}
	
	// TODO: logging
	
	// p := &progress.WriteCounter{}

	// kernelPath, sourceURL, destPath are global variables??
	//kernelPath, err := utils.DownloadFile(sourceURL, destPath, p)
	
	// return err
	return nil
}

