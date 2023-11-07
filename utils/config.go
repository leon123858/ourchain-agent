package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type OurChainConfig struct {
	ServerHost string `json:"serverHost"`
	ServerPort int    `json:"serverPort"`
	User       string `json:"user"`
	Passwd     string `json:"passwd"`
	UseSsl     bool   `json:"useSsl"`
	//WALLET_PASSPHRASE = "WalletPassphrase"
}

var OurChainConfigInstance OurChainConfig

func LoadConfig(path ...string) {
	if len(path) == 0 {
		path = make([]string, 1)
		path[0] = "./config.toml"
	}
	viper.SetConfigFile(path[0])

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	err = viper.Unmarshal(&OurChainConfigInstance)
	if err != nil {
		log.Fatalln(err)
	}
	// override config with env host (for production only)
	if os.Getenv("APP_HOST") != "" {
		OurChainConfigInstance.ServerHost = os.Getenv("APP_HOST")
	}
}
