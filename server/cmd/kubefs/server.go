package main

import (
	"sync"

	"github.com/cmwylie19/kubefs/pkg/server"
	"github.com/spf13/cobra"
)

var (

	// WaitGroup is used to wait for the program to finish goroutines.
	wg sync.WaitGroup
	// cfgPath is the path to the EnvoyGateway configuration file.
	key  string
	cert string
	port string
	dir  string
)

// getServerCommand returns the server cobra command to be executed.
func getServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Serve Media Controller",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := &server.Server{}
			return s.Serve(key, cert, dir, port)
			//		return s.Serve(tlsKey, tlsCert, port)
		},
	}

	cmd.PersistentFlags().StringVarP(&dir, "dir", "d", "/tmp", "FTP directory of the IP Camera")
	cmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "Port from which to server the webserver")
	cmd.PersistentFlags().StringVarP(&key, "key", "", "", "Server private key for TLS encryption.")
	cmd.PersistentFlags().StringVarP(&cert, "cert", "", "", "Server certificate for TLS encryption")
	return cmd
}
