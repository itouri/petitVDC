package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "petitvdc",
	Short: "",
	Long:  ``,
}

func Execute() {
	RootCmd.AddCommand(createCmd)
	if err := RootCmd.Execute; err != nil {
		fmt.Println(err)
		os.Exit(-1) // ??? -1?
	}
}
