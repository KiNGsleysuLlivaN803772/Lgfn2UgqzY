// 代码生成时间: 2025-09-17 02:52:33
package main

import (
    "fmt"
    "math"
    "os"

    "github.com/spf13/cobra"
)

// 定义版本变量
var version string

// 根命令
var rootCmd = &cobra.Command{
    Use:   "math-tool",
    Short: "math-tool is a mathematical calculation toolkit",
    Long:  "
math-tool is a mathematical calculation toolkit.
",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

// 添加 version 命令
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of math-tool",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("math-tool version", version)
    },
}

// 添加 add 子命令
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add two numbers",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Sum: %d
", addNumbers(args[0], args[1]))
    },
}

// 添加 subtract 子命令
var subtractCmd = &cobra.Command{
    Use:   "subtract",
    Short: "Subtract two numbers",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Difference: %d
", subtractNumbers(args[0], args[1]))
    },
}

// 添加 multiply 子命令
var multiplyCmd = &cobra.Command{
    Use:   "multiply",
    Short: "Multiply two numbers",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Product: %d
", multiplyNumbers(args[0], args[1]))
    },
}

// 添加 divide 子命令
var divideCmd = &cobra.Command{
    Use:   "divide",
    Short: "Divide two numbers",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Quotient: %d
", divideNumbers(args[0], args[1]))
    },
}

// 初始化 Cobra
func init() {
    rootCmd.AddCommand(versionCmd)
    rootCmd.AddCommand(addCmd, subtractCmd, multiplyCmd, divideCmd)
    cobra.OnInitialize(initConfig)

    // 配置命令行参数
    rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/math-tool.yaml)")
}

// 初始化配置
func initConfig() {
    // 读取配置文件
    if cfgFile != "" {
        fmt.Println("Using config file:", cfgFile)
        viper.SetConfigFile(cfgFile)
        if err := viper.ReadInConfig(); err != nil {
            cobra.CheckErr(err)
        }
    }
}

// addNumbers 添加两个整数
func addNumbers(a string, b string) int {
    aInt, _ := strconv.Atoi(a)
    bInt, _ := strconv.Atoi(b)
    return aInt + bInt
}

// subtractNumbers 减去两个整数
func subtractNumbers(a string, b string) int {
    aInt, _ := strconv.Atoi(a)
    bInt, _ := strconv.Atoi(b)
    return aInt - bInt
}

// multiplyNumbers 乘以两个整数
func multiplyNumbers(a string, b string) int {
    aInt, _ := strconv.Atoi(a)
    bInt, _ := strconv.Atoi(b)
    return aInt * bInt
}

// divideNumbers 除以两个整数
func divideNumbers(a string, b string) int {
    bInt, _ := strconv.Atoi(b)
    if bInt == 0 {
        fmt.Println("Error: Division by zero")
        os.Exit(1)
    }
    aInt, _ := strconv.Atoi(a)
    return aInt / bInt
}

// main 函数
func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, "Error: ", err)
        os.Exit(1)
    }
}
