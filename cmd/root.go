package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Special = "This is a special message."

var roorCmd = &cobra.Command{
	Use:   "go-cobra",
	Short: "A brief description of your application",
	Long:  `This is the long description of the command line application.`,
}

func Execute() {
	cobra.CheckErr(roorCmd.Execute())
}

func init() {
	roorCmd.PersistentFlags().StringP("directory", "d", os.TempDir(), "Path to use.")
	roorCmd.PersistentFlags().Uint("depth", 2, "depth to search.")

	viper.BindPFlag("directory", roorCmd.PersistentFlags().Lookup("directory"))
	viper.BindPFlag("depth", roorCmd.Flags().Lookup("depth"))
}
