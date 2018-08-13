package cmd

import (
	"github.com/spf13/cobra"
	"aws-autoscaling-tester/pkg/client"
	"aws-autoscaling-tester/pkg/client/config"
)

var conf config.Config

func NewCmdClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "client",
		Short:   "",
		Example: "",
		RunE:    runClient,
	}
	cmd.Flags().StringVarP(&conf.Url, "url", "u", "", "elb url with protocol (required)")
	cmd.Flags().Int32VarP(&conf.ConnectionValue, "conn", "c", 100, "connection value (required)")
	cmd.Flags().Int32VarP(&conf.Period, "period", "p", 30, "period second (required)")
	cmd.MarkFlagRequired("url")
	return cmd
}

func runClient(cmd *cobra.Command, args []string) error {
	c := client.NewClient(&conf)
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}
