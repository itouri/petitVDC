package cmd

import (
	"../../../api"
	"../internal/util"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create [ResourceTemplate.json]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return pflag.ErrHelp // ???
		}
		req := prepareCreateAPICall(args[0], args[1:], cmd.Flags())
	},
}

func prepareCreateAPICall(templateSlug string, args []string, flags *pflag.FlagSet) *api.CreateRequest, error {
	// パスを基にtemplate.jsonを見つける
	regTmpl, err := util.FetchTemplate(templateSlug)
	if err != nil {
		nil, err
	}



	// オプションに追加されたパラメータをjsonにくっつける
}
