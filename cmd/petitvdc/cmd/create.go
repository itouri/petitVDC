package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [ResourceTemplate.json]",
	Short: "Register an instance from template",
	Long:  `Register an instance from template.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// if len(args) < 1 {
		// 	return pflag.ErrHelp // ???
		// }
		// req := prepareCreateAPICall(args[0], args[1:], cmd.Flags())
		// call := func(conn *grpc.ClientConn) error {
		// 	c := api.NewInstanceClient(conn)
		// 	reply, err := c.Create(context.Background()), req) // req はクロージャ渡し
		// 	if err != nil {
		// 		log.WithError(err).Fatal("Disconnected abnormaly")
		// 		return err
		// 	}
		// 	fmt.Println(reply.GetInstanceId())
		// 	return err
		// }
		// return util.RemoteCall(call)
		fmt.Println("OK")
		return nil
	},
}

// func prepareCreateAPICall(templateSlug string, args []string, flags *pflag.FlagSet) *api.CreateRequest, error {
// 	// TODO URLで取ってるjsonとかに対応するのは後
// 	// パスを基にtemplate.jsonを見つける
// 	tmpl, err := util.FetchTemplate(templateSlug)
// 	if err != nil {
// 		nil, err
// 	}

// 	// TODO オプションに追加されたパラメータをjsonにくっつける

// 	req := &api.CreateRequest {
// 		InstanceTemplate: tmpl,
// 	}
// 	return req
// }
