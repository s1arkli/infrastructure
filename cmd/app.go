package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"architechture/app/api"
	"architechture/app/infrastructure/init/config"
)

func init() {
	rootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "启动http服务",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitViper()

		r := api.NewRouter()
		//todo just trust nginx
		r.SetTrustedProxies(nil)

		r.Run(":" + viper.GetString("port"))
	},
}
