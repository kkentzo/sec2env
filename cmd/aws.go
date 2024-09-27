package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

func AwsCommand(globals *GlobalVariables) *cobra.Command {
	var (
		// command-line arguments
		region     string
		secretName string
		descr      = "Access secret from AWS Secrets Manager"
		cmd        = &cobra.Command{
			Use:   "aws",
			Short: descr,
			Long:  descr,
			Run: func(cmd *cobra.Command, args []string) {
				svc := secretsmanager.New(session.New(), aws.NewConfig().WithRegion(region))
				input := &secretsmanager.GetSecretValueInput{
					SecretId:     aws.String(secretName),
					VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
				}
				result, err := svc.GetSecretValue(input)
				if err != nil {
					if globals.verbose {
						fmt.Printf("error retrieving the secret value: %v\n", err)
						return
					}
					os.Exit(1)
				}
				kv, err := parseSecretAsJSON(*result.SecretString)
				if err != nil {
					if globals.verbose {
						fmt.Printf("error parsing the secret's json contents: %v\n", err)
					}
					return
				}
				displayAsEnv(kv)
			},
		}
	)

	cmd.Flags().StringVarP(&region, "region", "r", "", "AWS region")
	cmd.Flags().StringVarP(&secretName, "name", "n", "", "secret name")

	cmd.MarkFlagRequired("region")
	cmd.MarkFlagRequired("name")

	return requireGlobalFlags(cmd, globals)
}
