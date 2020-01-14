package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	schemaregistry "github.com/WildBeavers/schema-registry"
)

var (
	cfgFile        string
	registryURL    string
	basicAuthUser  string
	basicAuthPass  string
	verbose        bool
	nocolor        bool
	noConfirmation bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "schema-registry-cli",
	Short: "A command line interface for the Confluent schema registry",
	Long:  `A command line interface for the Confluent schema registry`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !verbose {
			log.SetOutput(ioutil.Discard)
		}
		if nocolor {
			color.NoColor = true
		}
		log.Printf("schema registry url: %s\n", viper.Get("url"))
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "be verbose")
	RootCmd.PersistentFlags().BoolVarP(&nocolor, "no-color", "n", false, "dont color output")
	RootCmd.PersistentFlags().StringVarP(&registryURL, "url", "e", schemaregistry.DefaultURL, "schema registry url, overrides SCHEMA_REGISTRY_URL")
	RootCmd.PersistentFlags().StringVarP(&basicAuthUser, "basic-auth-user", "u", "", "User for basic auth, overrides SCHEMA_REGISTRY_BASIC_AUTH_USER")
	RootCmd.PersistentFlags().StringVarP(&basicAuthPass, "basic-auth-pass", "p", "", "Password for basic auth, overrides SCHEMA_REGISTRY_BASIC_AUTH_PASS")
	RootCmd.PersistentFlags().BoolVarP(&noConfirmation, "yes", "y", false, "skip confirmation prompt")
	viper.SetEnvPrefix("schema_registry")
	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
	viper.BindEnv("url")
	viper.BindPFlag("basic_auth_user", RootCmd.PersistentFlags().Lookup("basic-auth-user"))
	viper.BindEnv("basic_auth_user")
	viper.BindPFlag("basic_auth_pass", RootCmd.PersistentFlags().Lookup("basic-auth-pass"))
	viper.BindEnv("basic_auth_pass")
}
