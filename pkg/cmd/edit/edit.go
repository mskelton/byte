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
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hasArg := len(args) > 0
		bytes, err := storage.SearchBytes(func(byte storage.Byte) bool {
			// If an arg is provided, filter by id or slug
			if hasArg {
				return byte.Id == args[0] || byte.Slug == args[0]
			}

			return true
		})
		if err != nil {
			return err
		}

		if len(bytes) == 0 {
			fmt.Println("no byte found")
			return nil
		}

		var slug string

		// Auto-select the byte if there is only one
		if len(bytes) == 1 {
			slug = bytes[0].Slug
		} else {
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
			err = storage.SyncByte("Edit", filename)
			if err != nil {
				return err
			}
		}

		// Print the slug to allow filtering the output of this command
		fmt.Println(byte.Id)

		return nil
	},
}
