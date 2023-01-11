package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TODO
func RestMock() {
	logrus.Println("url:", viper.GetString("rest.url"))
}
