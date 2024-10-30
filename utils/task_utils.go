package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/feelthatvib3/go-task-tracker/models"
)

func ChangeTaskStatus(arg string, newStatus string) {
	id, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		fmt.Println("Invalid task ID:", arg)
		return
	}

	tasks, err := models.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	taskFound := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			taskFound = true
			break
		}
	}

	if !taskFound {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}

	if err := models.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
	} else {
		fmt.Printf("Task %d marked as %s successfully.\n", id, newStatus)
	}
}

func FormatDate(dateStr string) string {
	parsedDate, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return dateStr
	}
	return parsedDate.Format("02 Jan 2006 03:04 PM")
}
