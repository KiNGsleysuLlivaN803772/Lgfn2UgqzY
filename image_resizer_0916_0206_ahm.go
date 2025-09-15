// 代码生成时间: 2025-09-16 02:06:50
package main

import (
    "flag"
    "fmt"
    "image"
    "image/jpeg"
    "os"
    "path/filepath"
    "log"
    "strings"

    "github.com/nfnt/resize"
)

// Define the command line flags
var (
    srcDir  = flag.String("src", "./source", "Directory containing source images")
    dstDir  = flag.String("dst\, "./destination", "Directory to store resized images")
    width   = flag.Int("width", 1024, "Width of the resized images")
    height  = flag.Int("height\, 768", "Height of the resized images")
)

// Function to resize an image
func resizeImage(srcPath, dstPath string, width, height int) error {
    srcFile, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    img, _, err := image.Decode(srcFile)
    if err != nil {
        return err
    }

    // Resize the image
    resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

    // Create the destination file
    dstFile, err := os.Create(dstPath)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    // Save the resized image
    err = jpeg.Encode(dstFile, resizedImg, nil)
    if err != nil {
        return err
    }

    return nil
}

// Function to process all images in a directory
func processDirectory(srcDir, dstDir string, width, height int) error {
    if err := os.MkdirAll(dstDir, 0755); err != nil {
        return err
    }

    dir, err := os.ReadDir(srcDir)
    if err != nil {
        return err
    }

    for _, entry := range dir {
        if !entry.IsDir() {
            srcPath := filepath.Join(srcDir, entry.Name())
            baseName := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
            dstPath := filepath.Join(dstDir, baseName+"_resized"+filepath.Ext(entry.Name()))
            if err := resizeImage(srcPath, dstPath, width, height); err != nil {
                log.Printf("Failed to resize image %s: %v", srcPath, err)
            }
        }
    }

    return nil
}

func main() {
    flag.Parse()

    if err := processDirectory(*srcDir, *dstDir, *width, *height); err != nil {
        fmt.Printf("Error processing images: %v\
", err)
    } else {
        fmt.Println("Images processed successfully")
    }
}