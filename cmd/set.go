package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("set called host=%s, port=%d\n", host, port)

		for _, arg := range args {
			fmt.Printf("%s\n", arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
