package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var confItems = map[string][]string {
	"bot": {"token", "debug"},
}

func Init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Configuration file loaded")

	for k, v := range confItems {
		checkConfIsSet(k, v)
	}

	logrus.Info("Configuration file checking succeeded")
}

func checkConfIsSet(name string, keys []string) {
	for i := range keys {
		wholeKey := name + "." + keys[i]
		if !viper.IsSet(wholeKey) {
			logrus.WithField(wholeKey, nil).
				Fatal("The following item of your configuration file hasn't been set properly: ")
		}
	}
	return
}
