// 代码生成时间: 2025-09-21 20:31:07
package main

import (
    "fmt"
    "math"
    "os"

    "github.com/spf13/cobra"
)

// 定义一个函数来执行加法运算
func add(args []float64) float64 {
    return args[0] + args[1]
}

// 定义一个函数来执行减法运算
func subtract(args []float64) float64 {
    return args[0] - args[1]
}

// 定义一个函数来执行乘法运算
func multiply(args []float64) float64 {
    return args[0] * args[1]
}

// 定义一个函数来执行除法运算
func divide(args []float64) float64 {
    if args[1] == 0 {
        return math.NaN() // 防止除以零的错误
    }
    return args[0] / args[1]
}

// rootCmd 代表应用的根命令
var rootCmd = &cobra.Command{
    Use:   "math-tool",
    Short: "A brief description of your application",
    Long: `
A longer description of your application.
For example:
math-tool add 5 3
`,
    Args: cobra.ExactArgs(2), // 确保恰好有两个参数
    Run: func(cmd *cobra.Command, args []string) {
        // 将输入参数转换为float64类型
        num1, num2 := 0.0, 0.0
        if _, err := fmt.Sscanf(args[0], "%f", &num1); err != nil {
            fmt.Println("Error: Invalid input for the first number.")
            os.Exit(1)
        }

        if _, err := fmt.Sscanf(args[1], "%f", &num2); err != nil {
            fmt.Println("Error: Invalid input for the second number.")
            os.Exit(1)
        }

        // 根据操作符执行相应的数学运算
        switch cmd.Flag("operation").Value.String() {
        case "add":
            fmt.Println("Result: ", add([]float64{num1, num2}))
        case "subtract":
            fmt.Println("Result: ", subtract([]float64{num1, num2}))
        case "multiply":
            fmt.Println("Result: ", multiply([]float64{num1, num2}))
        case "divide":
            result := divide([]float64{num1, num2})
            if math.IsNaN(result) {
                fmt.Println("Error: Division by zero is not allowed.")
                os.Exit(1)
            }
            fmt.Println("Result: ", result)
        default:
            fmt.Println("Error: Unknown operation.")
            os.Exit(1)
        }
    },
}

// init 函数用于初始化命令行参数
func init() {
    // 添加一个名为"operation"的标记，用于指定操作类型
    rootCmd.PersistentFlags().StringP(
        "operation",
        "o",
        "",
        "The operation to perform (add, subtract, multiply, divide)",
    )
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}