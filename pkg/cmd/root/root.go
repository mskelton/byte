package root

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mskelton/byte/internal/editor"
	"github.com/mskelton/byte/internal/storage"
	"github.com/mskelton/byte/internal/utils"
	"github.com/mskelton/byte/pkg/cmd/id"
	"github.com/mskelton/byte/pkg/cmd/list"
	"github.com/mskelton/byte/pkg/cmd/search"
	"github.com/mskelton/byte/pkg/cmd/tag"
	"github.com/mskelton/byte/pkg/cmd/url"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "byte",
	Short: "Create a new byte",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Pull the latest bytes from the repo if the user specified the --commit
		// flag. By the time the file is saved, this should be done.
		if viper.GetBool("commit") {
			go storage.Pull()
		}

		// Create a temporary file
		tempFilename, err := editor.CreateTempFile("byte-*.md", []byte(TEMPLATE))
		if err != nil {
			return err
		}

		editor := editor.Editor{
			Editor:   viper.GetString("editor"),
			Filename: tempFilename,
		}

		// Open the users editor to edit the file
		data, err := editor.Edit()
		if err != nil {
			return err
		}

		// Remove the temp file
		defer func() {
			os.Remove(editor.Filename)
		}()

		// Bail if the user didn't enter any content
		if len(bytes.TrimSpace(data)) <= len(strings.TrimSpace(TEMPLATE)) {
			fmt.Println("byte content was empty, canceling operation")
			return nil
		}

		id := utils.GenerateId()
		filename, err := storage.GetBytePath(id)
		if err != nil {
			return err
		}

		// Write the byte to the file system
		err = storage.WriteByte(filename, data)
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
		fmt.Println(id)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/byte/config.yaml)")
	rootCmd.PersistentFlags().String("dir", "", "directory where bytes are stored")
	rootCmd.Flags().BoolP("commit", "c", false, "commit new bytes to Git after saving")
	rootCmd.Flags().String("editor", "", "editor to use when creating a new byte")

	viper.BindPFlag("commit", rootCmd.Flags().Lookup("commit"))
	viper.BindPFlag("dir", rootCmd.Flags().Lookup("dir"))
	viper.BindPFlag("editor", rootCmd.Flags().Lookup("editor"))

	rootCmd.AddCommand(id.IdCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(url.UrlCmd)
	rootCmd.AddCommand(tag.TagListCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(path.Join(home, ".config", "byte"))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; use defaults
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
