// 代码生成时间: 2025-09-30 22:36:16
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
)

// Task represents a single task
type Task struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Assignee  string `json:"assignee"`
    Completed bool   `json:"completed"`
}

// Tasks is a collection of tasks
type Tasks []Task

// AddTaskCmd represents the add task command
var AddTaskCmd = &cobra.Command{
    Use:   "add [title]",
    Short: "Add a new task",
    Args:  cobra.MinimumNArgs(1),
    Run:   addTask,
}

// ListTaskCmd represents the list tasks command
var ListTaskCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run:   listTasks,
}

// tasksData holds the collection of tasks
var tasksData Tasks

func main() {
    var rootCmd = &cobra.Command{
        Use:   "task-manager",
        Short: "A brief description of your application",
    }

    rootCmd.AddCommand(AddTaskCmd)
    rootCmd.AddCommand(ListTaskCmd)

    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

// addTask adds a new task to the collection
func addTask(cmd *cobra.Command, args []string) {
    title := args[0]
    task := Task{ID: fmt.Sprintf("%v", len(tasksData)+1), Title: title, Completed: false}
    tasksData = append(tasksData, task)
    fmt.Println("Task added:", task.ID)
}

// listTasks lists all tasks
func listTasks(cmd *cobra.Command, args []string) {
    if len(tasksData) == 0 {
        fmt.Println("No tasks found.")
        return
    }
    fmt.Println("Tasks:")
    for _, task := range tasksData {
        fmt.Printf("ID: %s, Title: %s, Assignee: %s, Completed: %v
", task.ID, task.Title, task.Assignee, task.Completed)
    }
}

// saveTasks saves tasks to a file
func saveTasks() {
    data, err := json.MarshalIndent(tasksData, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    if err := os.WriteFile("tasks.json", data, 0644); err != nil {
        log.Fatal(err)
    }
}

// loadTasks loads tasks from a file
func loadTasks() {
    if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
        return
    }
    data, err := os.ReadFile("tasks.json")
    if err != nil {
        log.Fatal(err)
    }
    if err := json.Unmarshal(data, &tasksData); err != nil {
        log.Fatal(err)
    }
}

func init() {
    loadTasks()
    cobra.OnInitialize(saveTasks)
}