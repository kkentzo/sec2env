package cmd

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/spf13/cobra"
)

func AzureCommand(globals *GlobalVariables) *cobra.Command {
	var (
		// command-line arguments
		vaultName  string
		secretName string
		descr      = "Access secret from Azure Vault"
		cmd        = &cobra.Command{
			Use:   "azure",
			Short: descr,
			Long:  descr,
			Run: func(cmd *cobra.Command, args []string) {
				vaultURI := fmt.Sprintf("https://%s.vault.azure.net/", vaultName)
				// Create a credential using the NewDefaultAzureCredential type.
				cred, err := azidentity.NewDefaultAzureCredential(nil)
				if err != nil {
					fmt.Printf("failed to obtain a credential: %v", err)
					return
				}

				// Establish a connection to the Key Vault client
				client, err := azsecrets.NewClient(vaultURI, cred, nil)
				if err != nil {
					fmt.Printf("failed to create the client: %v", err)
					return
				}

				// Get a secret. An empty string version gets the latest version of the secret.
				resp, err := client.GetSecret(context.Background(), secretName, "", nil)
				if err != nil {
					fmt.Printf("failed to get the secret: %v", err)
					return
				}

				kv, err := parseSecretAsJSON(*resp.Value)
				if err != nil {
					if globals.verbose {
						fmt.Printf("error parsing the secret's json contents: %v\n", err)
						return
					}
					return
				}
				displayAsEnv(kv)
			},
		}
	)

	cmd.Flags().StringVarP(&vaultName, "vault", "v", "", "vault name")
	cmd.Flags().StringVarP(&secretName, "name", "n", "", "secret name")

	cmd.MarkFlagRequired("vault")
	cmd.MarkFlagRequired("name")

	return requireGlobalFlags(cmd, globals)
}
