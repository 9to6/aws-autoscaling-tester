package cmd

import (
	"github.com/spf13/cobra"
	"github.com/9to6/cloudwatch-metrics-tester/pkg/server"
)


func NewCmdServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Short:   "",
		Example: "",
		RunE:    runServer,
	}
	return cmd
}

func runServer(cmd *cobra.Command, args []string) error {
	server, err := server.NewServer()
	if err != nil {
		return err
	}
	if err := server.Run(); err != nil {
		return err
	}
	return nil
}
