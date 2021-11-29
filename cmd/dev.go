package cmd

import (
	"github.com/razvvan/CryptoYay-Server/internal/yay"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Dev Commands",
}

var sendSampleYayCmd = &cobra.Command{
	Use: "send-yay",
	Run: func(cmd *cobra.Command, args []string) {
		err := yay.Send("0x70997970c51812dc3a010c7d01b50e0d17dc79c8", "NoOrg")
		cobra.CheckErr(err)
	},
}

func init() {
	devCmd.AddCommand(sendSampleYayCmd)
	rootCmd.AddCommand(devCmd)
}
