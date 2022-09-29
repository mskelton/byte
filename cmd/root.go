package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/mskelton/zet/cmd/id"
	"github.com/mskelton/zet/cmd/last"
	"github.com/mskelton/zet/util"
	"github.com/mskelton/zet/util/editor"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "zet",
	Short: "Create a new zettel",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the repo from the config and throw an error if it is not specified
		repo := viper.GetString("repo")
		if repo == "" {
			return errors.New("no zettel repo defined. Please define one in $HOME/.config/zet/config.yaml.")
		}

		// Resolve the path to the repo that was specified in the config
		resolvedRepo, err := util.ResolvePath(repo)
		if err != nil {
			return err
		}

		id := util.Id()
		editor := editor.Editor{
			Editor:   viper.GetString("editor"),
			Filename: path.Join(resolvedRepo, id, "README.md"),
		}

		// Create the directory for the zettel if necessary
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
			fmt.Println("zettel content was empty, canceling operation")

			// Remove the empty zettel and it's parent directory
			err := os.RemoveAll(path.Dir(editor.Filename))
			if err != nil {
				return err
			}

			return nil
		}

		// Print the id to allow filtering the output of this command
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/zet/config.yaml)")
	rootCmd.PersistentFlags().String("repo", "", "repository where zettels are stored")
	rootCmd.Flags().BoolP("commit", "c", false, "commit new zettels to VCS after saving")
	rootCmd.Flags().String("editor", "", "editor to use when creating a new zettel")
	viper.BindPFlag("commit", rootCmd.Flags().Lookup("commit"))
	viper.BindPFlag("editor", rootCmd.Flags().Lookup("editor"))
	viper.BindPFlag("repo", rootCmd.Flags().Lookup("repo"))

	rootCmd.AddCommand(id.IdCmd)
	rootCmd.AddCommand(last.LastCmd)
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

		viper.AddConfigPath(path.Join(home, ".config", "zet"))
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
