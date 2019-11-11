package cmd

import (
	"log"

	"github.com/galo/moloon/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "start the moloon controller",
	Long:  `Starts a http controller server and serves the api. `,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer(true)
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(controllerCmd)

	// Here you will define your flags and configuration settings.
	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("log_level", "debug")

	// Agents - by default use kubernetes discovery

	// Dev mode
	viper.SetDefault("dev", "true")

	viper.SetDefault("auth_jwk_url", "https://core.api.hp.com/hpid/oauth/v1/jwks")
	viper.SetDefault("auth_jwt_secret", "random")
	viper.SetDefault("auth_jwt_expiry", "15m")
	viper.SetDefault("auth_jwt_refresh_expiry", "1h")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
