package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var confItems = map[string][]string {
	"bot": {"token", "debug"},
	"forward": {"src", "dest"},
}

type forwardEnd struct {
	UserNameArray []string
	IdArray       []int64
}

var ForwardSrc forwardEnd
var ForwardDest forwardEnd

func InitConf() {
	viper.SetConfigName("conf")  // set the config file name. Viper will automatically detect the file extension name
	viper.AddConfigPath("./")     // search the config file under the current directory

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error while reading config file: %s \n", err))
	}

	fmt.Println("Configuration file loaded.")

	for k, v := range confItems {
		err := checkConfIsSet(k, v)
		if err {
			panic(fmt.Sprintf("\"%s\" item of your config file hasn't been set properly. \nPlease check your config file.", k))
		}
	}

	ForwardSrc.classifyForwardConfig(viper.Get("forward.src"))
	ForwardDest.classifyForwardConfig(viper.Get("forward.dest"))

	fmt.Println("Configuration file checking succeeded. All required values are set.")
}

func checkConfIsSet(name string, keys []string) (getKeyErrExists bool) {
	getKeyErrExists = false
	for i := range keys {
		if !viper.IsSet(name + "." + keys[i]) {
			fmt.Printf("\"%s\" not set", keys[i])
			fmt.Printf("in\"" + name + "\"\n")
			getKeyErrExists = true
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
	case string:
		f.UserNameArray = append(f.UserNameArray, t)
	default:
		panic(`"forward" field of the config file hasn't been set correctly`)
	}
}
