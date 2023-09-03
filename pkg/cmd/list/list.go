package list

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/byte/internal/storage"
	"github.com/spf13/cobra"
)

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
		filename, err := storage.GetByteDir()
		if err != nil {
			return err
		}

		// List all the files in the byte dir
		files, err := os.ReadDir(filename)
		if err != nil {
			return err
		}

		if len(files) == 0 {
			fmt.Println("no bytes found")
			return nil
		}

		// Sort the files by name
		sort.Slice(files, func(i, j int) bool {
			return strings.Compare(files[i].Name(), files[j].Name()) > 0
		})

		for i := len(files) - 1; i >= 0; i-- {
			name := files[i].Name()
			slug := strings.TrimSuffix(name, ".md")

			fmt.Println(slug)
		}

		return nil
	},
}

func init() {
	ListCmd.Aliases = []string{"ls"}
}
