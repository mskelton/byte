package cmd

import (
	"os"

	"github.com/mskelton/zet/cmd/id"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zet",
	Short: "TODO",
	Long:  `TODO`,
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(id.IdCmd)
}
