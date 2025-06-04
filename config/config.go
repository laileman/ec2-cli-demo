package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Aws AWSConfig
}

type AWSConfig struct {
	AccessKeyID string
	SecretKeyID string
	Region      string
}

func NewConfig(config string) *Config {
	viper.SetConfigName(config)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	// 使用嵌套 key 获取 aws 配置
	Region := viper.GetString("aws.region")
	AccessKey := viper.GetString("aws.access_key_id")
	SecretKey := viper.GetString("aws.secret_access_key")

	return &Config{
		Aws: AWSConfig{
			AccessKeyID: AccessKey,
			SecretKeyID: SecretKey,
			Region:      Region,
		},
	}
}
