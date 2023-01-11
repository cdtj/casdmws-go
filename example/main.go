package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Setting highest loggin level
	logrus.SetLevel(logrus.TraceLevel)
	// passing file name...
	// private.yaml is ignored by .gitignore,
	// check with sample.yaml
	viper.SetConfigName("private")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// printing config
	viper.Debug()
	// SOAP Mock Sample
	SoapMock()
	// SOAP Prod Sample
	SoapProd()
	// REST Prod Sample
	RestProd()
	// REST Mock Sample
	RestMock()
	// REST Prod Attmnt Upload
	RestAttmnt()
}
