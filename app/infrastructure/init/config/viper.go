package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cwd, _ := os.Getwd()
	cfgPath := filepath.Join(cwd, "config.yaml")

	//<ConfigPath>/<ConfigName>.<ConfigType>
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cwd) // 加上这行

	check(cfgPath)
}

func check(cfgPath string) {
	err := viper.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			viper.Set("port", 9595)

			//config not found
			//create config file
			if err = viper.WriteConfigAs(cfgPath); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			//failed to parse config file
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
