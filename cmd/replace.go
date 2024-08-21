package cmd

import (
	"github.com/luka2220/tools/ccmc/internal/app/client"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SendReplaceCommand(host, port, args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
}
