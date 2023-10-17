package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/byte/internal/storage"
	"github.com/spf13/cobra"
)

var printSlug bool

type ListOptions struct {
	slug bool
}

func ListBytes(options ListOptions) ([]string, error) {
	filename, err := storage.GetByteDir()
	if err != nil {
		return []string{}, err
	}

	// List all the files in the byte dir
	files, err := os.ReadDir(filename)
	if err != nil {
		return []string{}, err
	}

	if len(files) == 0 {
		fmt.Println("no bytes found")
		return []string{}, nil
	}

	ids := make([]string, len(files))
	for i := len(files) - 1; i >= 0; i-- {
		ids[i] = strings.TrimSuffix(files[i].Name(), ".md")
	}

	return ids, nil
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved bytes",
	Long: heredoc.Doc(`
    This command will list all saved bytes in descending order showing the most
    recently added bytes first. This command accepts a few arguments for 
    filtering and sorting, but advanced filtering and sorting should be done
    using the 'byte find' command.
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

		for _, byte := range bytes {
			if printSlug {
				fmt.Println(byte.Slug)
			} else {
				fmt.Println(byte.Id)
			}
		}

		return nil
	},
}

func init() {
	ListCmd.Aliases = []string{"ls"}
	ListCmd.Flags().BoolVarP(&printSlug, "slug", "s", false, "Show slugs instead of ids")
}
