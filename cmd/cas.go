package cmd

import (
	"fmt"
	"strconv"

	"github.com/luka2220/tools/ccmc/internal/app/client"
	"github.com/spf13/cobra"
)

var casCmd = &cobra.Command{
	Use:   "cas",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error, the last argument (token) must be a number")
			return
		}

		client.SendCasCommand(host, port, args[0], args[1], token)
	},
}

func init() {
	rootCmd.AddCommand(casCmd)
}
