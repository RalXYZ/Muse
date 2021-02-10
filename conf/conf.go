package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
)

var confItems = map[string][]string {
	"bot": {"token", "debug"},
	"forward": {"src", "dest"},
}

type forwardEnd struct {
	IdArray       []int64
}

var ForwardSrc forwardEnd
var ForwardDest forwardEnd

func Init() {
	viper.SetConfigName("conf")  // set the config file name. Viper will automatically detect the file extension name
	viper.AddConfigPath("./")     // search the config file under the current directory

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Configuration file loaded")

	for k, v := range confItems {
		checkConfIsSet(k, v)
	}

	ForwardSrc.classifyForwardConfig(viper.Get("forward.src"))
	ForwardDest.classifyForwardConfig(viper.Get("forward.dest"))

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

func (f *forwardEnd) classifyForwardConfig(conf interface{}) {
	if array, ok := conf.([]interface{}); ok {
		for _, v := range array {
			f.setForwardValue(v)
		}
		return
	}
	f.setForwardValue(conf)
}

func (f *forwardEnd) setForwardValue(conf interface{}) {
	switch t := conf.(type) {
	case int:
		f.IdArray = append(f.IdArray, int64(t))
	default:
		logrus.WithField(fmt.Sprintf(`TypeOf(%v)`, t), reflect.TypeOf(t)).
			Fatal(`Configuration file hasn't been set correctly, type error:`)
	}
}
