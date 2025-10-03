// 代码生成时间: 2025-10-04 02:58:23
package main

import (
    "flag"
    "fmt"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "strings"

    // 导入cobra库
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "batch_rename_tool",
    Short: "A tool for batch renaming files",
    Long:  `A simple tool for batch renaming files based on certain patterns or rules.`,
    Run: func(cmd *cobra.Command, args []string) {
        pattern, _ := cmd.Flags().GetString("pattern")
        renameFiles(pattern)
    },
}

// renameFiles 执行批量重命名操作
func renameFiles(pattern string) {
    files, err := os.ReadDir(".")
    if err != nil {
        log.Fatalf("Failed to read directory: %v", err)
    }
    for _, file := range files {
        if !file.IsDir() {
            oldPath := file.Name()
            newPath := generateNewName(oldPath, pattern)
            if err := os.Rename(oldPath, newPath); err != nil {
                log.Printf("Failed to rename %s to %s: %v", oldPath, newPath, err)
                continue
            }
            fmt.Printf("Renamed %s to %s", oldPath, newPath)
        }
    }
}

// generateNewName 根据给定的模式生成新文件名
func generateNewName(oldName, pattern string) string {
    parts := strings.Split(oldName, ".")
    if len(parts) == 0 {
        return oldName
    }
    baseName := strings.Join(parts[:len(parts)-1], ".")
    extension := parts[len(parts)-1]
    return fmt.Sprintf("%s_%s.%s", baseName, pattern, extension)
}

func main() {
    // 添加命令行参数
    patternFlag := rootCmd.Flags().StringP("pattern", "p", "new", "The pattern to append to the filename")
    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Error executing command: %v", err)
    }
}
