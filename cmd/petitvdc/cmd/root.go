package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var RootCmd = &cobra.Command{
	Use:   "petitvdc",
	Short: "",
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	RootCmd.AddCommand(createCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1) // ??? -1?
	}
}
