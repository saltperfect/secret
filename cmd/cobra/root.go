package cobra

import (
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
)

var RootCmd = &cobra.Command{
	Use:   "secret",
	Short: "Secret is API key storage and secret manager",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the encoding key" )
}

func secretPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secret")
}
