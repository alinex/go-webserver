// Package cmd contains CLI commands called through cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgFile holds an alternative config file path given as cli flag
var cfgFile string

// init will initialize the cobra parsing.
func init() {
	// initialize config through viper
	cobra.OnInitialize(initConfig)
	// Define the flags to be supported.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-webserver.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-webserver")   // name of config file (without extension)
	viper.AddConfigPath(os.Getenv("HOME")) // adding home directory as first search path
	viper.AutomaticEnv()                   // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	// define flags and use config file value before the above default
	RootCmd.Flags().IntP("port", "p", 8080, "port to listen for http requests")
	viper.BindPFlag("port", RootCmd.Flags().Lookup("port"))
	// append commands
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
