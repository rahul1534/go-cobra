package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Special = "This is a special message."

var rootCmd = &cobra.Command{
	Use:   "go-cobra",
	Short: "A brief description of your application",
	Long:  `This is the long description of the command line application.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("directory", "d", os.TempDir(), "Path to use.")
	rootCmd.PersistentFlags().Uint("depth", 2, "depth to search.")

	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))
	viper.BindPFlag("depth", rootCmd.Flags().Lookup("depth"))
}
