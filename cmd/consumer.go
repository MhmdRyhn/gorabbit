package cmd

import (
	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "This command will start the consumer for the RabbitMQ.",
	Run:   consumer,
}

func consumer(cmd *cobra.Command, args []string) {}
