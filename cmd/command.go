package cmd

import (
	"github.com/spf13/cobra"
)

var mainCommand = &cobra.Command{
	Use:   "gorabbit",
	Short: "gorabbit sample application for RabbitMQ",
}

func Execute() {
	mainCommand.AddCommand(producerCmd)
	mainCommand.AddCommand(consumerCmd)
	mainCommand.Execute()
}
