package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"kman/pkg/progress"
)

func ConfirmAction(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + " (Y/n) ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return false
	}

	input = strings.TrimSpace(input)
	return strings.ToLower(input) == "y" || input == ""
}

func DownloadFile(url, destPath string, p progress.Progress) (string, error) {
	// Get file name and create full path
	filePath := filepath.Join(destPath, filepath.Base(url))

	// Ensure destination directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return "", fmt.Errorf("error creating directory: %w", err)
	}

	// If file already exists then return
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("File already exists: %s\n", filePath)
		return filePath, nil
	}

	// Open file for writing
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %w", err)
	}
	defer outFile.Close()

	// Get the file from the URL
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error downloading file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	// Initialize progress tracking
	p.Start(resp.ContentLength)

	// Write response body to file
	_, err = io.Copy(outFile, io.TeeReader(resp.Body, p.(*progress.WriteCounter)))
	if err != nil {
		return "", fmt.Errorf("error writing file: %w", err)
	}

	p.Finish()
	return filePath, nil
}

func UncompressFile(filePath, extractDir string) error {
	ext := filepath.Ext(filePath)

	var cmd *exec.Cmd

	if err := os.MkdirAll(extractDir, 0755); err != nil {
		return fmt.Errorf("failed to create extraction directory: %v", err)
	}

	switch ext {
	case ".gz":
		cmd = exec.Command("tar", "-xzf", filePath, "-C", extractDir)
	case ".xz":
		cmd = exec.Command("tar", "-xJf", filePath, "-C", extractDir)
	default:
		return fmt.Errorf("unsupported file extension: %s", ext)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Uncompressing: %s\n", filePath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to uncompress file: %v", err)
	}

	return nil
}

func IsPackageInstalled(pkg string) bool {
	_, err := exec.LookPath(pkg)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func ShowSpinner(done chan bool) {
	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	for {
		select {
		case <-done:
			return
		default:
			for _, s := range spinner {
				fmt.Printf("\r%s", s)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
