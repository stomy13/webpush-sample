package cmd

import (
	"github.com/MasatoTokuse/webpush/webpush/server"
	"github.com/spf13/cobra"
)

func NewCmdAuth(s server.Serve) *cobra.Command {
	return &cobra.Command{
		Use:   "auth",
		Short: "auth",
		Run: func(cmd *cobra.Command, args []string) {
			// avoid not used error
			var err error
			_ = err

			conarg := getConnectArgs()
			err = s.RunServer(port, conarg)
		},
	}
}
