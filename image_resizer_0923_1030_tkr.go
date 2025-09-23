// 代码生成时间: 2025-09-23 10:30:45
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
    "golang.org/x/image/bmp"
    "golang.org/x/image/tiff"
# FIXME: 处理边界情况
    "image"
# NOTE: 重要实现细节
    "image/jpeg"
    "image/png"
)

// version is the current version of the application.
var version = "0.1.0"

// resizeOptions contains the options for resizing images.
type resizeOptions struct {
    width  int
    height int
# 增强安全性
    folder string
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
    Use:   "image-resizer [options]",
    Short: "A brief description of your application",
# 扩展功能模块
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This app is a tool to resize images.`,

    // Run is the main entry point when the command is executed.
    Run: func(cmd *cobra.Command, args []string) {
        // Initialize resize options.
        options := resizeOptions{
            width:  cmd.Int("width"),
            height: cmd.Int("height\),
            folder: cmd.String("folder\),
        }

        // Validate options.
        if options.width <= 0 || options.height <= 0 {
            fmt.Println("Error: Width and height must be positive integers.")
# FIXME: 处理边界情况
            return
        }

        // Recursively process all files in the specified folder.
        if err := resizeImagesInFolder(options.folder, options.width, options.height); err != nil {
            log.Fatalf("Failed to resize images: %v", err)
        }
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
# 扩展功能模块
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
# 增强安全性
        os.Exit(1)
    }
}

// resizeImagesInFolder resizes all images in the specified folder.
func resizeImagesInFolder(folder string, width, height int) error {
    files, err := os.ReadDir(folder)
    if err != nil {
# 扩展功能模块
        return err
    }
# TODO: 优化性能

    for _, file := range files {
        if !file.IsDir() {
            filePath := filepath.Join(folder, file.Name())
            if err := resizeImage(filePath, width, height); err != nil {
                fmt.Printf("Failed to resize image %s: %v", filePath, err)
            }
        }
    }

    return nil
# 优化算法效率
}

// resizeImage resizes the image at the specified path.
func resizeImage(filePath string, width, height int) error {
    src, err := os.Open(filePath)
    if err != nil {
# NOTE: 重要实现细节
        return err
    }
    defer src.Close()

    img, _, err := image.Decode(src)
# FIXME: 处理边界情况
    if err != nil {
        return err
    }
# 改进用户体验

    // Create a new image with the desired dimensions.
    newImg := image.NewRGBA(image.Rect(0, 0, width, height))
    newImg = resize.Resize(width, height, img, resize.Lanczos3)

    // Save the resized image.
# TODO: 优化性能
    outputFileName := fmt.Sprintf("%s_resized%s", filepath.Base(filePath), filepath.Ext(filePath))
    out, err := os.Create(filepath.Join(filepath.Dir(filePath), outputFileName))
    if err != nil {
        return err
    }
    defer out.Close()

    // Determine the image format and encode to the output file.
    ext := filepath.Ext(filePath)
    switch ext {
    case ".jpg":
        return jpeg.Encode(out, newImg, nil)
    case ".png\):
        return png.Encode(out, newImg)
    case ".bmp":
        return bmp.Encode(out, newImg)
    case ".tiff":
# NOTE: 重要实现细节
        return tiff.Encode(out, newImg)
    default:
        return fmt.Errorf("unsupported image format: %s", ext)
    }
}

func init() {
    // Define flags.
    rootCmd.Flags().IntP("width", "w", 0, "Width for resizing images.")
    rootCmd.Flags().IntP("height", "h", 0, "Height for resizing images.")
    rootCmd.Flags().StringP("folder", "f", ".", "Folder containing images to resize.")
# 改进用户体验
    rootCmd.Flags().BoolP("version", "v", false, "Print the version of the application.")

    // Set the version.
    rootCmd.VersionTemplate = `Version: {{.Version}}`
    rootCmd.SetVersionTemplate("{{.Version}}
\)
    rootCmd.Version = version

    // Mark the flags as required.
    rootCmd.MarkFlagRequired("width\)
    rootCmd.MarkFlagRequired("height\)
    rootCmd.MarkFlagRequired("folder\)
}
# FIXME: 处理边界情况

func main() {
# 增强安全性
    Execute()
}
# TODO: 优化性能