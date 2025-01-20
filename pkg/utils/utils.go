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
	"errors"

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

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DownloadFile(url, destPath string, p progress.Progress) (string, error) {
	// Get file name and create full path
	result := filepath.Join(destPath, filepath.Base(url))
	result, err := filepath.Abs(result)
	if err != nil {
		return "", fmt.Errorf("error resolving absolute path: %w", err)
	}

	// Ensure destination directory exists
	if err := os.MkdirAll(filepath.Dir(result), 0755); err != nil {
		return "", fmt.Errorf("error creating directory: %w", err)
	}

	if FileExists(result) {
		fmt.Printf("File already exists: %s\n", result)
		return result, nil
	}

	// Open file for writing
	outFile, err := os.Create(result)
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
	return result, nil
}

func UncompressFile(archivePath, extractDir string) ([]string, error) {
	extension := filepath.Ext(archivePath)
	var cmd *exec.Cmd

	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("archive file does not exist: %s", archivePath)
	}

	if err := os.MkdirAll(extractDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create extraction directory: %v", err)
	}

	beforeFiles, err := listFiles(extractDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list files before extraction: %w", err)
	}

	switch extension {
	case ".gz":
		cmd = exec.Command("tar", "-xzf", archivePath, "-C", extractDir)
	case ".xz":
		cmd = exec.Command("tar", "-xJf", archivePath, "-C", extractDir)
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", extension)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Uncompressing: %s\n", archivePath)

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run command: %v", err)
	}

	afterFiles, err := listFiles(extractDir)
	if err != nil {
		return nil, fmt.Errorf("failed to list files before extraction: %w", err)
	}

	extractedFiles := findNewFiles(beforeFiles, afterFiles)
	if len(extractedFiles) == 0 {
		return nil, fmt.Errorf("no files or folders extracted from archive")
	}
	
	return extractedFiles, nil
}

func listFiles(dir string) (map[string]struct{}, error) {
	files := make(map[string]struct{})
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files[path] = struct{}{}
		}
		return nil
	})
	return files, err
}

func findNewFiles(before, after map[string]struct{}) []string {
	newFiles := []string{}
	for file := range after {
		if _, exists := before[file]; !exists {
			newFiles = append(newFiles, file)
		}
	}
	return newFiles
}

func RemoveFile(filePath string) error {
	info, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	if info.IsDir() {
		return fmt.Errorf("path is a directory, not a file: %s", filePath)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to remove file: %s, error: %w", filePath, err)
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

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}

