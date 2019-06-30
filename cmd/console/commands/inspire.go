package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inspireCmd = &cobra.Command{
	Use:   "inspire",
	Short: "Print an inspiring quote",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You learn more from failure than from success.")
	},
}

func init() {
	rootCmd.AddCommand(inspireCmd)
}
