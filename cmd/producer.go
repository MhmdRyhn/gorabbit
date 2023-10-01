package cmd

import (
	"github.com/spf13/cobra"
)

var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "This command will start the producer for the RabbitMQ.",
	Run:   producer,
}

func producer(cmd *cobra.Command, args []string) {}
