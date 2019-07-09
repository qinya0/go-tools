package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
)

var RootCmd = &cobra.Command{
	Use: "go-tools",
	Short: "a tools write by golang",
	Version: version,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(pwdCmd)
	RootCmd.AddCommand(testCmd)
}