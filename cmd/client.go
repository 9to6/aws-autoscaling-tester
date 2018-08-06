package cmd

import "github.com/spf13/cobra"

func NewCmdClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "client",
		Short:   "",
		Example: "",
		RunE:    runClient,
	}
	return cmd
}

func runClient(cmd *cobra.Command, args []string) error {
	return nil
}
