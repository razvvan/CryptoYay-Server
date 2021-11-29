package cmd

import (
	"github.com/razvvan/CryptoYay-Server/internal/serve"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Listen",
	Run: func(cmd *cobra.Command, args []string) {
		err := serve.Listen()
		if err != nil {
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
