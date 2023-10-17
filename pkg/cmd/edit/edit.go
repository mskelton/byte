package edit

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/byte/internal/utils"
	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a byte",
	Long: heredoc.Doc(`
    Edits a byte given a byte id or slug. If no id or slug is provided, a fuzzy
    search will be performed to find the byte to edit.
  `),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(utils.GenerateId())
	},
}
