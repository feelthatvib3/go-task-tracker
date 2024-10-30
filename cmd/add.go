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

	"github.com/feelthatvib3/go-task-tracker/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task DESCRIPTION]",
	Short: "Adds a new task",
	Long:  `Adds a new task to the task list with the status set to "todo".`,

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		tasks, err := models.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		var maxID int64 = 0
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		newID := maxID + 1

		newTask := models.NewTask(newID, description)

		tasks = append(tasks, newTask)
		if err := models.SaveTasks(tasks); err != nil {
			fmt.Println("Error saving task:", err)
		} else {
			fmt.Printf("Task added successfully: %s (ID: %d)\n", description, newID)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
