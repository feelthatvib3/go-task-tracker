/*
Copyright © 2024 Vladyslav Skrynnyk

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/feelthatvib3/go-task-tracker/models"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [task ID] [new description]",
	Short: "Update the description of an existing task",
	Long:  `Updates the description of a specified task by its ID.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		newDescription := args[1]

		tasks, err := models.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		taskFound := false
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Description = newDescription
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
			fmt.Printf("Task %d updated successfully.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
