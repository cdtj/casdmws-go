package main

import (
	"github.com/cdtj/casdmws-go/soap"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SoapProd is a sample where we are using login by:
//   - user/password;
//   - policy and cert;
//
// and:
//   - creating a new request with type "I" (Incident)
func SoapProd() {
	logrus.Println("url:", viper.GetString("soap.url"))

	// Test run w/ Login/Password auth
	pc1 := SoapSample{
		soap.NewSdmSoap(viper.GetString("soap.url")),
		soap.NewLoginUserPassword(viper.GetString("soap.username"),
			viper.GetString("soap.password"),
			viper.GetString("soap.policy"),
		)}
	if err := pc1.run(); err != nil {
		logrus.Errorln("prod cli 1 failed with: ", err)
	}

	// Test run w/ certificate auth
	pc2 := SoapSample{
		soap.NewSdmSoap(viper.GetString("soap.url")),
		soap.NewLoginManaged(viper.GetString("soap.policy"),
			viper.GetString("soap.cert"),
		)}

	if err := pc2.run(); err != nil {
		logrus.Errorln("prod cli 2 failed with: ", err)
	}

}
