package cmd

import (
	"github.com/9to6/aws-autoscaling-tester/pkg/server"
	"github.com/9to6/aws-autoscaling-tester/pkg/server/config"
	"github.com/spf13/cobra"
)

var serverConf config.Config

func NewCmdServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Short:   "",
		Example: "",
		RunE:    runServer,
	}
	cmd.Flags().Int32VarP(&serverConf.ConnectionValue, "conn", "c", 100, "default connection value")
	cmd.Flags().IntVarP(&serverConf.Port, "port", "", 8080, "REST API port")
	return cmd
}

func runServer(cmd *cobra.Command, args []string) error {
	s, err := server.NewServer(serverConf)
	if err != nil {
		return err
	}
	if err := s.Run(); err != nil {
		return err
	}
	return nil
}
