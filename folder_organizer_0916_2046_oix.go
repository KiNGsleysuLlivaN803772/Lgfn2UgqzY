// 代码生成时间: 2025-09-16 20:46:54
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// FolderOrganizer is a program that organizes files in a directory
type FolderOrganizer struct {
    srcDir string
}

// NewFolderOrganizer creates a new FolderOrganizer instance
func NewFolderOrganizer(srcDir string) *FolderOrganizer {
    return &FolderOrganizer{
        srcDir: srcDir,
    }
}

// Organize organizes files in the source directory
func (f *FolderOrganizer) Organize() error {
    // Check if the source directory exists
    if _, err := os.Stat(f.srcDir); os.IsNotExist(err) {
        return fmt.Errorf("source directory does not exist: %s", f.srcDir)
    }

    // Get a list of files in the source directory
    files, err := os.ReadDir(f.srcDir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %s", err)
    }

    // Iterate through the files and organize them
    for _, file := range files {
        if file.IsDir() {
            continue // Skip directories
        }

        // Get the file extension
        ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")

        // Create a new directory for the file type
        dirName := fmt.Sprintf("%s/%s_files", f.srcDir, ext)
        if _, err := os.Stat(dirName); os.IsNotExist(err) {
            if err := os.MkdirAll(dirName, 0755); err != nil {
                return fmt.Errorf("failed to create directory: %s", err)
            }
        }

        // Move the file to the new directory
        srcPath := filepath.Join(f.srcDir, file.Name())
        dstPath := filepath.Join(dirName, file.Name())
        if err := os.Rename(srcPath, dstPath); err != nil {
            return fmt.Errorf("failed to move file: %s", err)
        }
    }

    return nil
}

func main() {
    // Define flags
    srcDirFlag := flag.StringP("srcDir", "s", ".", "source directory to organize")

    // Parse flags
    flag.Parse()

    // Create a new FolderOrganizer instance
    organizer := NewFolderOrganizer(*srcDirFlag)

    // Organize files in the source directory
    if err := organizer.Organize(); err != nil {
        log.Fatalf("error organizing files: %s", err)
    }

    fmt.Println("Files organized successfully")
}
