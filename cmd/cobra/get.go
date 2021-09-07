package cobra

import (
	"fmt"
	"github.com/saltperfect/exercise/secret"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command {
	Use: "get",
	Short: "Get a secret in your secret manager",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.NewVault(encodingKey,secretPath())
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			fmt.Println("no value set")
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

var getAllCmd = &cobra.Command{
	Use: "getAll",
	Short: "Get all stored values",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.NewVault(encodingKey,secretPath())
		value, err := v.GetAll()
		if err != nil {
			fmt.Println("no value set")
		}
		fmt.Printf("%+v\n", value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	RootCmd.AddCommand(getAllCmd)
}