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

	exe, _ := os.Executable()
	base := filepath.Dir(exe)
	cfgPath := filepath.Join(base, "config.yaml")
	//<ConfigPath>/<ConfigName>.<ConfigType>
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(base)

	check(cfgPath)
}

func check(cfgPath string) {
	if err := viper.ReadInConfig(); err != nil {
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
