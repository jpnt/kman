package progress

import (
	"fmt"
)

// Interface for tracking progress
type Progress interface {
	Start(total int64)
	Update(current int64)
	Finish()
}

type WriteCounter struct {
	Total   int64
	Current int64
	Unknown bool
}

func (wc *WriteCounter) Start(total int64) {
	wc.Total = total
	wc.Current = 0
	wc.Unknown = total <= 0 // Mark unkown if total is negative
	if wc.Unknown {
		fmt.Println("Download started... (unknown total size)")
	} else {
		fmt.Println("Download started...")
	}
}

func (wc *WriteCounter) Update(current int64) {
	wc.Current += current
	if wc.Unknown {
		fmt.Printf("\rDownloaded %d bytes", wc.Current)
	} else {
		fmt.Printf("\rProgress: %.2f%% (%d/%d bytes)",
		float64(wc.Current)/float64(wc.Total)*100, wc.Total, wc.Current)
	}
}

func (wc *WriteCounter) Finish() {
	fmt.Println("\nDownload complete.")
}

// Create a proxy for io.Writer to capture progress updates
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Update(int64(n))
	return n, nil
}
