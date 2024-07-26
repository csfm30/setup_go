package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	initConfig()
	fmt.Println("Hello World")
}

func initConfig() {
	switch os.Getenv("ENV") {
	case "":
		os.Setenv("ENV", "dev")
		viper.SetConfigName("config_dev")
	default:
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
