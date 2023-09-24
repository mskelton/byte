package url

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func printUrl(slug string) {

	fmt.Println("https://mskelton.dev/bytes/" + slug)
}

var UrlCmd = &cobra.Command{
	Use:   "url",
	Short: "Print the byte URL given the slug.",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// If an arg was provided, print the URL for that slug
		if len(args) == 1 {
			printUrl(args[0])
			return nil
		}

		// Otherwise, read slugs from stdin and print the URLs
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			printUrl(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		return nil
	},
}
