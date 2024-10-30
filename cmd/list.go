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
	"os"

	"github.com/feelthatvib3/go-task-tracker/models"
	"github.com/feelthatvib3/go-task-tracker/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List all tasks or filter by status",
	Long:  `Lists all tasks. If a status is provided, only tasks with that status will be displayed.`,

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var statusFilter string
		if len(args) > 0 {
			statusFilter = args[0]
		}

		tasks, err := models.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Description", "Status", "Created At", "Updated At"})

		taskFound := false

		for _, task := range tasks {
			if statusFilter == "" || task.Status == statusFilter {
				table.Append([]string{
					fmt.Sprintf("%d", task.ID),
					task.Description,
					task.Status,
					utils.FormatDate(task.CreatedAt),
					utils.FormatDate(task.UpdatedAt),
				})
				taskFound = true
			}
		}

		if taskFound {
			table.Render()
		} else {
			fmt.Printf("No tasks found with status '%s'.\n", statusFilter)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
