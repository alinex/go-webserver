// Package cmd contains CLI commands called through cobra.
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-webserver [-p <port>]",
	Short: "Simple Web Server",
	Long:  `It will serve the current directory through the web server.`,
	// Running the root command
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("port")
		fmt.Printf("Server started under port %v...\n", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), http.FileServer(http.Dir(".")))
	},
}
