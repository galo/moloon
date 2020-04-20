package cmd

import (
	"log"

	"github.com/galo/moloon/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// agentCmd represents the serve command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "start the moloon agent",
	Long:  `starts a http agent and serves the api`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := api.NewServer(false)
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// In development you can use "localhost:3000" to avoid issues with local firewealls
	viper.SetDefault("port", ":3000")
	viper.SetDefault("log_level", "debug")

	// Dev mode
	viper.SetDefault("dev", "true")

	viper.SetDefault("auth_jwk_url", "https://core.api.hp.com/hpid/oauth/v1/jwks")
	viper.SetDefault("auth_jwt_secret", "random")
	viper.SetDefault("auth_jwt_expiry", "15m")
	viper.SetDefault("auth_jwt_refresh_expiry", "1h")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
