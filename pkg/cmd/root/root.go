package root

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/mskelton/byte/internal/editor"
	"github.com/mskelton/byte/internal/storage"
	"github.com/mskelton/byte/pkg/cmd/list"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "byte",
	Short: "Create a new byte",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slug := args[0]
		filename, err := storage.GetBytePath(slug)
		if err != nil {
			return err
		}

		editor := editor.Editor{
			Editor:   viper.GetString("editor"),
			Filename: filename,
		}

		// Create the directory for the bytetel if necessary
		err = os.Mkdir(path.Dir(editor.Filename), os.ModePerm)
		if err != nil && !(errors.Is(err, os.ErrExist)) {
			return err
		}

		// Open the users editor to edit the file
		data, err := editor.Edit()
		if err != nil {
			return err
		}

		// Fail if the user didn't enter any content
		if len(data) == 0 {
			fmt.Println("bytetel content was empty, canceling operation")

			// Remove the empty byte and it's parent directory
			err := os.Remove(editor.Filename)
			if err != nil {
				return err
			}

			return nil
		}

		// Print the slug to allow filtering the output of this command
		fmt.Println(slug)
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

	rootCmd.AddCommand(list.ListCmd)
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
