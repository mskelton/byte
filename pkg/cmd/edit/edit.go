package edit

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/byte/internal/storage"
	"github.com/mskelton/byte/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a byte",
	Long: heredoc.Doc(`
    Edits a byte given a byte id or slug. If no id or slug is provided, a fuzzy
    search will be performed to find the byte to edit.
  `),
	RunE: func(cmd *cobra.Command, args []string) error {
		bytes, err := storage.SearchBytes(nil)
		if err != nil {
			return err
		}

		if len(bytes) == 0 {
			fmt.Println("no bytes found")
			return nil
		}

		var slugs []string
		for _, item := range bytes {
			slugs = append(slugs, item.Slug)
		}

		slug, err := utils.Filter(slugs)
		if err != nil {
			return err
		}
		if slug == "" {
			fmt.Println("no bytes found")
		}

		var byte storage.Byte
		for _, b := range bytes {
			if b.Slug == slug {
				byte = b
			}
		}

		filename, err := storage.EditByte(byte.Id)
		if err != nil {
			return err
		}

		// Commit and push the byte to Git if the user specified the --commit flag
		if viper.GetBool("commit") {
			err = storage.SyncByte(filename)
			if err != nil {
				return err
			}
		}

		// Print the slug to allow filtering the output of this command
		fmt.Println(byte.Id)

		return nil
	},
}
