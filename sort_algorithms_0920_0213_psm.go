// 代码生成时间: 2025-09-20 02:13:08
 * a command-line interface to select and run different sorting algorithms.
 *
 * Usage:
 *   go run sort_algorithms.go --help
 */

package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/spf13/cobra"
)

// Sortable is an interface that any sortable type must implement.
type Sortable interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// IntSlice attaches the methods of Sortable to []int,
// sorting in increasing order.
type IntSlice []int

func (p IntSlice) Len() int { return len(p) }

func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// BubbleSort sorts a slice of ints using bubble sort algorithm.
func BubbleSort(data Sortable) {
    for i := 0; i < data.Len(); i++ {
        for j := 0; j < data.Len()-i-1; j++ {
            if data.Less(j+1, j) {
                data.Swap(j, j+1)
            }
        }
    }
}

// QuickSort sorts a slice of ints using quicksort algorithm.
func QuickSort(data Sortable) {
    if data.Len() < 2 {
        return
    }
    var pivot int = data.Len()/2
    var low, high int = 0, data.Len()-1
    var mid int
    for low <= high {
        for low <= high && data.Less(high, pivot) {
            high--
        }
        data.Swap(high, low)
        for low <= high && data.Less(pivot, low) {
            low++
        }
        data.Swap(low, high)
        if pivot == low {
            mid = low
        } else {
            mid = high
        }
        high = mid - 1
        low = mid + 1
    }
    if mid-pivot < 0 {
        QuickSort(data)
    } else if mid-pivot > 0 {
        QuickSort(data.Slice(0, mid))
    }
    if data.Len()-(mid+1)-pivot < 0 {
        QuickSort(data.Slice(mid+1, data.Len()))
    } else {
        QuickSort(data.Slice(mid+1, data.Len()))
    }
}

// CmdSort is the command for sorting algorithms.
var CmdSort = &cobra.Command{
    Use:   "sort",
    Short: "Run sorting algorithms",
    Long:  "Run sorting algorithms (bubble sort and quicksort) from the command line.",
    Run: func(cmd *cobra.Command, args []string) {
        randomData := generateRandomIntSlice(10)
        fmt.Println("Original data: ", randomData)
        fmt.Println("Sorting with Bubble Sort...")
        BubbleSort(randomData)
        fmt.Println("Sorted data (Bubble Sort): ", randomData)
        fmt.Println("Sorting with Quick Sort...")
        QuickSort(randomData)
        fmt.Println("Sorted data (Quick Sort): ", randomData)
    },
}

func main() {
    var rootCmd = &cobra.Command{
        Use:   "sort_algorithms",
        Short: "Sort Algorithms CLI",
        Long:  "Sort Algorithms CLI is a tool for demonstrating sorting algorithms.",
    }
    rootCmd.AddCommand(CmdSort)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}

// generateRandomIntSlice generates a slice of random integers.
func generateRandomIntSlice(size int) IntSlice {
    rand.Seed(time.Now().UnixNano())
    randomData := make(IntSlice, size)
    for i := 0; i < size; i++ {
        randomData[i] = rand.Intn(100)
    }
    return randomData
}