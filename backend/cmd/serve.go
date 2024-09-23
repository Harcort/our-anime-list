/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"our-anime-list/backend/internal/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		server.SetValues(host, port)
		server.StartServer()
	},
}

var host string
var port string

func init() {
	serveCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "This flag sets the host of our API server")
	serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This flag sets the port of our API server")

	rootCmd.AddCommand(serveCmd)

}
