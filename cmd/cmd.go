package cmd

import (
	"log"
)

var taskTitle string
var all, comp, uncomp bool

func Exec() {
	addCmd.Flags().StringVarP(&taskTitle, "title", "t", "", "task title flag")

	readCmd.Flags().BoolVarP(&comp,"completed", "c", false, "for reading just completed tasks")
	readCmd.Flags().BoolVarP(&uncomp,"uncompleted", "u", false, "for reading just uncompleted tasks")
	readCmd.Flags().BoolVarP(&all,"all", "a", false, "for reading all tasks")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(readCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error on executing root command: %s", err)
	}
}