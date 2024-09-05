package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/luka2220/tools/ccmc/internal/app/client"
	"github.com/spf13/cobra"
)

var (
	host string
	port int
)

var rootCmd = &cobra.Command{
	Use:   "ccmc",
	Short: "Client for interacting with a memcached server",
	Long: `Welcome to the memcahced CLI!

	Memcached is a free, open-source, high-performance, distributed memory object caching system. It is intended to speed up dynamic web applications by reducing database load. It is also an in-memory key-value store for small chunks of arbitrary data retrieved from back-end systems with higher latency. It is simple yet powerful. Its simple design promotes quick deployment, and ease of development, and solves many problems facing large data caches. Its relatively simple API is available for most popular languages. It uses a simple text-based network protocol, making it a great platform to learn how to build network clients and servers.

	The CLI client communicates with memcached over a TCP connection. The support commands in this client are:

	set - store some data
	add - store some data, only if the server doesn't already hold data for this key
	replace - store some data, only if the server already holds data for this key
	append - Add the data to an existing key, after the existing data
	prepend - Add the data to an exisiting key, before the existing data
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hostname=%s port=%d", host, port)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SendDeleteCommand(host, port, args[0])
	},
}

var incrementCmd = &cobra.Command{
	Use:   "incr",
	Args:  cobra.ExactArgs(2),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		value, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error, the increment value must be an integer")
			return
		}

		client.SendIncrCommand(host, port, args[0], value)
	},
}

var decrementCmd = &cobra.Command{
	Use:   "decr",
	Args:  cobra.ExactArgs(2),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		value, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error, the decrement value must be an integer")
			return
		}

		client.SendDecrCommand(host, port, args[0], value)
	},
}

var appendCmd = &cobra.Command{
	Use:   "append",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SendAppendCommand(host, port, args[0], args[1])
	},
}

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
		client.SendSetCommand(host, port, args[0], args[1])
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SendGetCommand(host, port, args[0])
	},
}

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

var prependCmd = &cobra.Command{
	Use:   "prepend",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SendPrependCommmand(host, port, args[0], args[1])
	},
}

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

var getsCmd = &cobra.Command{
	Use:   "gets",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SendGetsCommand(host, port, args[0])
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.SendAddCommand(host, port, args[0], args[1])
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&host, "host", "o", "localhost", "Host address of tcp client")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 11211, "Port of the client address")

	// Commands
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(incrementCmd)
	rootCmd.AddCommand(decrementCmd)
	rootCmd.AddCommand(appendCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(replaceCmd)
	rootCmd.AddCommand(prependCmd)
	rootCmd.AddCommand(casCmd)
	rootCmd.AddCommand(getsCmd)
	rootCmd.AddCommand(addCmd)
}
