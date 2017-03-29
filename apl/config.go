package apl

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
)

// Config holds the values gathered from the environment
type Config struct {
	API      string
	Username string
	Password string
}

var APLConfig Config

func init() {
	err := ProcessConfigs()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// ProcessConfigs ...
func ProcessConfigs() error {

	// Check to see if already processed
	if (Config{}) != APLConfig {
		return nil
	}

	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("$HOME/.apl")
	v.AddConfigPath(".")

	// Try to read from conf, if not found ignore
	v.ReadInConfig()


	v.SetEnvPrefix("apl")
	v.BindEnv("api")
	v.BindEnv("svc_username")
	v.BindEnv("svc_password")

	api := v.GetString("api")
	username := v.GetString("svc_username")
	password := v.GetString("svc_password")

	if api == "" || username == "" || password == "" {
		var errorString string
		if api == ""{
			errorString += "Missing env var \"APL_API\" or config file value \"api\".\n"
		}

		if username == "" {
			errorString += "Missing env var \"APL_SVC_USERNAME\" or config file value \"svc_username\".\n"
		}

		if password == "" {
			errorString += "Missing env var \"APL_SVC_USERNAME\" or config file value \"svc_username\".\n"
		}

		return fmt.Errorf(errorString)
	}

	APLConfig = Config{
		API: api,
		Username: username,
		Password: password,
	}
	//fmt.Println(APLConfig)
	return nil
}

