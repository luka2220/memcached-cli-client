package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "memcached-client",
	Short: "Client for interacting with a memcached server",
	Long:  `Memcached is a simple yet powerful server with client libraries in many programming languages. But sometimes it’s useful to be able to interact with a server from the command line. For example it’s useful to be able to use curl to fetch web pages or test/automate calls to RESTful API.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
