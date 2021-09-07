package cobra

import (
	"fmt"

	"github.com/saltperfect/exercise/secret"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set a secret in your secret manager",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.NewVault(encodingKey, secretPath())
		key := args[0]
		value := args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
