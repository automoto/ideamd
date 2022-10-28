package pkg

import (
    "github.com/spf13/viper"
    "path/filepath"
)

func ParseConfig(configPath string) error {
    cn := filepath.Base(configPath)
    viper.SetConfigType("yaml")
    viper.SetConfigName(cn)
    viper.AddConfigPath(configPath)
    viper.AddConfigPath(".")
    viper.AutomaticEnv()
    viper.SetDefault("doc.batch", "False")
    err := viper.ReadInConfig()
    if err != nil {
        return err
    } else {
        return nil
    }
}
