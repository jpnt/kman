package kernel

import "fmt"

type DownloadCommand struct {}

func (c *DownloadCommand) Execute() error {
	fmt.Println("DOWNLOAD COMMAND SUCCESS!!!")
	return nil
}

