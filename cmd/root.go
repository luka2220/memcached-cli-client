package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccmc",
	Short: "Client for interacting with a memcached server",
	Long: `Welcome to eh memcahced CLI!

	Memcached is a free, open-source, high-performance, distributed memory object caching system. It is intended to speed up dynamic web applications by reducing database load. It is also an in-memory key-value store for small chunks of arbitrary data retrieved from back-end systems with higher latency. It is simple yet powerful. Its simple design promotes quick deployment, and ease of development, and solves many problems facing large data caches. Its relatively simple API is available for most popular languages. It uses a simple text-based network protocol, making it a great platform to learn how to build network clients and servers.

	The CLI client communicates with memcached over a TCP connection. The support commands in this client are:

	set - store some data
	add - store some data, only if the server doesn't already hold data for this key
	replace - store some data, only if the server already holds data for this key
	append - Add the data to an existing key, after the existing data
	prepend - Add the data to an exisiting key, before the existing data
	`,
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
