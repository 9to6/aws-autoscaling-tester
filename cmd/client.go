package cmd

import (
	"fmt"
	"github.com/9to6/aws-autoscaling-tester/pkg/client"
	"github.com/9to6/aws-autoscaling-tester/pkg/client/config"
	"github.com/spf13/cobra"
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
	cmd.Flags().BoolVarP(&conf.IncMode, "inc", "", false, "increase mode")
	cmd.Flags().BoolVarP(&conf.DecMode, "dec", "", false, "decrease mode")
	cmd.Flags().Int32VarP(&conf.IncValue, "inc-val", "", 1, "increase connection value (default 1)")
	cmd.Flags().Int32VarP(&conf.DecValue, "dec-val", "", 1, "decrease connection value (default 1)")
	cmd.Flags().IntVarP(&conf.Port, "port", "", 8081, "REST API port")
	cmd.MarkFlagRequired("url")
	return cmd
}

func runClient(cmd *cobra.Command, args []string) error {
	if conf.IncMode && conf.DecMode {
		return fmt.Errorf("incMode and decMode cannot be set both")
	}
	c := client.NewClient(&conf)
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}
