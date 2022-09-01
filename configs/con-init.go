package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nqd/flat"
	"github.com/spf13/viper"
)

var config Config

func GetConfig() *Config {
	return &config
}

func ConfigsInit() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	env := (viper.Get("APP_ENV"))

	switch env {
	case "local":
		getEnv()
	case "develop":
		getEnv()
	default:
		getEnv()
	}
}

func getEnv() {
	secretData := map[string]interface{}{}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		secretData[pair[0]] = os.Getenv(pair[0])
	}

	addEnv, err := flat.Unflatten(lower(secretData), &flat.Options{
		Delimiter: "_",
	})

	j, _ := json.Marshal(addEnv)
	fmt.Println(string(j))

	if err != nil {
		panic(err.Error())
	}

	jsonConfig, err := json.Marshal(addEnv)
	if err != nil {
		fmt.Println(err)
	}

	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.ReadConfig(bytes.NewReader(jsonConfig))
	viper.Unmarshal(&config)

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
	}
}

func lower(v map[string]interface{}) map[string]interface{} {

	lv := make(map[string]interface{}, len(v))
	for mk, mv := range v {
		lv[strings.ToLower(mk)] = mv
	}
	return lv
}
