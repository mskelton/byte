package search

import (
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/byte/internal/storage"
	"github.com/spf13/cobra"
)

var tag string
var printSlug bool
var fullText bool

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for bytes",
	Long: heredoc.Doc(`
    Search for bytes by title, tag, or content. This does a search on the entire
    byte content, so it's a bit more powerful than the website's search.
  `),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Search for matching bytes
		bytes, err := storage.SearchBytes(func(byte storage.Byte) bool {
			if tag != "" {
				return byte.HasTag(tag)
			}

			// Collect all args into a single string
			query := strings.Join(args, " ")

			// First try searching by title
			if strings.Contains(strings.ToLower(byte.Title), strings.ToLower(query)) {
				return true
			}

			// If full text search is enabled, search by content
			if fullText {
				return strings.Contains(byte.Content, query)
			}

			return false
		})

		if err != nil {
			return err
		}

		// Print matching files
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
	SearchCmd.Flags().StringVarP(&tag, "tag", "t", "", "Search by tag")
	SearchCmd.Flags().BoolVarP(&fullText, "full-text", "f", false, "Full text search")
	SearchCmd.Flags().BoolVarP(&printSlug, "slug", "s", false, "Show slugs instead of ids")
}
