package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "wire",
	Short: "An application to demonstrate new project structure",
	Long:  "An application to demonstrate new project structure",
	Example: "go run main.go start-http OR \n" +
		"wire start-http",
	TraverseChildren: true,
	Hidden:           true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
