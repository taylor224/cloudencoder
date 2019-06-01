package cmd

import (
	"fmt"
	"runtime"
	"github.com/spf13/cobra"
	"github.com/alfg/enc/api/config"
	"github.com/alfg/enc/api/server"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Start the server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server...")
		startServer()
	},
}

func configRuntime() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Running with %d CPUs\n", numCPU)
}
func startServer() {
	// Get workflow configs.
	// // helpers.LoadConfig()
	// fmt.Println(helpers.C)

	// Create HTTP Server.
	configRuntime()
	port := config.GetPort()
	server.NewServer(port)
}
