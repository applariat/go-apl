package apl

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Config holds the values gathered from the environment
type Config struct {
	API          string `default:"https://api.applariat.io/v1/"`
	Svc_Username string `required:"true"`
	Svc_Password string `required:"true"`
}

var conf Config

func init() {
	err := envconfig.Process("apl", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

}
