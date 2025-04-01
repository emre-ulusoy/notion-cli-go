package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"notioncli/utils"
)

var addCmd = &cobra.Command{
	Use:   "add <task>",
	Short: "Add a new db page",
	Long:  `Add a new db page to the brain dump db`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		notionAPIKey, DBID := utils.SetAPIConfig()
		result := utils.AddNewDatabasePage(notionAPIKey, DBID, text)
		if result != nil {
			fmt.Printf("Error adding new task: %s\n", result)
			os.Exit(1)
		}
		fmt.Printf("Task %s added.\n", text)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	checkCmd.Flags().String("text", "", "Text for the new task")
}
