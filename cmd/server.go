package cmd

import (
	"github.com/spf13/cobra"
	"aws-autoscaling-tester/pkg/server"
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
	s, err := server.NewServer()
	if err != nil {
		return err
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
