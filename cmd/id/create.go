package id

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/zet/util"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new zet id",
	Long: heredoc.Doc(`
    Create a new zet id.

    In most cases you do not need to use this command directly as ids are
    automatically created for you. However, if you need to create a new zet id
    for an external program, this can be useful.
  `),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(util.Id())
	},
}
