package cmd

import (
	"fmt"

	"github.com/alfg/enc/api/config"
	"github.com/alfg/enc/api/worker"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start the worker.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting worker...")
		startWorkers()
	},
}

func startWorkers() {

	// Worker config.
	workerCfg := &worker.Config{
		Host:        config.Get().RedisHost,
		Port:        config.Get().RedisPort,
		Namespace:   config.Get().WorkerNamespace,
		JobName:     config.Get().WorkerJobName,
		Concurrency: config.Get().WorkerConcurrency,
	}

	// Create Workers.
	worker.NewWorker(*workerCfg)
}
