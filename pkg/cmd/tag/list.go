package tag

import (
	"sort"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/byte/internal/storage"
	"github.com/spf13/cobra"
)

var TagListCmd = &cobra.Command{
	Use:   "tag list",
	Short: "List all tags",
	Long: heredoc.Doc(`
    List all tags found in all bytes.
  `),
	RunE: func(cmd *cobra.Command, args []string) error {
		bytes, err := storage.GetAllBytes()
		if err != nil {
			return err
		}

		// Get a list of unique tag names
		tags := make(map[string]bool)
		for _, b := range bytes {
			for _, t := range b.Tags {
				tags[t] = true
			}
		}

		// Sort
		sortedTags := make([]string, 0, len(tags))
		for t := range tags {
			sortedTags = append(sortedTags, t)
		}
		sort.Strings(sortedTags)

		// Print the tags
		for _, tag := range sortedTags {
			cmd.Println(tag)
		}

		return nil
	},
}
