package core

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

const defaultEnv = "dev"

// LoadConfig ...
func LoadConfig(path string) error {

	env, ok := os.LookupEnv("APP_ENV")
	if !ok {
		env = defaultEnv
	}
	viper.AddConfigPath(path)
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	keys := viper.AllKeys()
	sort.Strings(keys)
	for _, k := range keys {
		v := viper.GetString(k)
		fmt.Println(k, "=", v)
	}

	return nil
}
