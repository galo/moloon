package cmd

import (
	"log"

	"github.com/galo/moloon/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "start the moloon server",
	Long:  `Starts a http server and serves the api`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer(false)
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.
	viper.SetDefault("port", "localhost:3000")
	viper.SetDefault("log_level", "debug")

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
