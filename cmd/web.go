package cmd

import (
	"github.com/spf13/cobra"
	"go-wire-architecture/cmd/bootstrap"
)

var webCmd = &cobra.Command{
	Use:   "start-http",
	Short: "A http go wire application",
	Long:  "A go wire application to demonstrate new project structure with http",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.WebApplication()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
