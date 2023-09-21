package id

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/byte/internal/utils"
	"github.com/spf13/cobra"
)

var IdCmd = &cobra.Command{
	Use:   "id",
	Short: "Generate a new byte id",
	Long: heredoc.Doc(`
    Bytes are stored as markdown using a unique id representing the date and time
    the byte was created. This command will generate a new id and print it to
    stdout. This is typically not needed as the id is automatically generated
    when creating a new byte.
  `),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(utils.GenerateId())
	},
}
