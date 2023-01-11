package main

import (
	"github.com/cdtj/casdmws-go/rest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// RestAttmnt is a sample where we are using login by user/password and:
//   - Uploading test png image using UploadFile function;
//   - Uploading test string wrapped into []byte using Upload,
//     and downloading it into file using DownloadFile;
//   - Logging out.
func RestAttmnt() {
	// Login with Login/Password
	cli1 := RestSample{rest.NewSdmRest(
		viper.GetString("rest.url"),
		viper.GetString("rest.repo"),
		viper.GetString("rest.server"),
		rest.NewSchema(viper.GetString("rest.schema")),
		rest.NewLoginUserPassword(
			viper.GetString("rest.username"),
			viper.GetString("rest.password"),
		),
		nil,
	)}
	if err := cli1.upload(); err != nil {
		logrus.WithFields(logrus.Fields{
			"cli": cli1,
			"err": err,
		}).Error("upload failed")
	}
	if err := cli1.uploadAndDownload(); err != nil {
		logrus.WithFields(logrus.Fields{
			"cli": cli1,
			"err": err,
		}).Error("uploadAndDownload failed")
	}
}

// RestProd is a sample where we are using login by:
//   - Secret Key;
//   - Username and Password;
//
// and:
//   - Fetchig CallReq (Request) status by it's REL_ATTR (code);
//   - Logging out.
func RestProd() {
	logrus.Println("url:", viper.GetString("rest.url"))

	// Login with Secret Key
	cli1 := RestSample{rest.NewSdmRest(
		viper.GetString("rest.url"),
		viper.GetString("rest.repo"),
		viper.GetString("rest.server"),
		rest.NewSchema(viper.GetString("rest.schema")),
		rest.NewLoginSecretKey(
			viper.GetString("rest.username"),
			viper.GetString("rest.password"),
			nil, // nil for uninstalled or NoValue [NX_STRING_TO_SIGN_FIELDS]
			rest.HmacAlgorithm(viper.GetUint("rest.hmac")), // stored in [NX_HMAC_ALGORITHM]
		),
		nil,
	)}
	if err := cli1.run(); err != nil {
		logrus.WithFields(logrus.Fields{
			"cli": cli1,
			"err": err,
		}).Error("failed")
	}

	// Login with Username and Password
	cli2 := RestSample{(rest.NewSdmRest(
		viper.GetString("rest.url"),
		viper.GetString("rest.repo"),
		viper.GetString("rest.server"),
		rest.NewSchema(viper.GetString("rest.schema")),
		rest.NewLoginUserPassword(
			viper.GetString("rest.username"),
			viper.GetString("rest.password"),
		),
		nil,
	))}
	if err := cli2.run(); err != nil {
		logrus.WithFields(logrus.Fields{
			"cli": cli2,
			"err": err,
		}).Error("failed")
	}
}
