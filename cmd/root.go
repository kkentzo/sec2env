package cmd

import (
	"github.com/spf13/cobra"
)

type GlobalVariables struct {
	verbose bool
}

func New() *cobra.Command {
	globals := &GlobalVariables{}

	descr := "Access AWS or Azure cloud secrets"

	root := &cobra.Command{
		Use:   "sec2env",
		Short: descr,
		Long:  descr,
	}
	root.AddCommand(AwsCommand(globals))
	root.AddCommand(AzureCommand(globals))
	return root
}

func requireGlobalFlags(cmd *cobra.Command, globals *GlobalVariables) *cobra.Command {
	cmd.Flags().BoolVar(&globals.verbose, "verbose", false, "verbose command output")
	return cmd
}
