package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
)

var FwdRuleMap = make(map[int64] []int64)

func LoadRuleMap() {
	rawRuleMap := viper.GetStringMap("rule")
	for rawMapKey, srcID := range rawRuleMap {
		arr, ok := srcID.([]interface{})
		if !ok {
			typeError(srcID)
		}

		int64MapKey, err := strconv.ParseInt(rawMapKey, 10, 64)
		if err != nil {
			typeError(rawMapKey)
		}

		FwdRuleMap[int64MapKey] = make([]int64, len(arr))
		for arrKey, rawDestID := range arr {
			intDestID, ok := rawDestID.(int)
			if !ok {
				typeError(rawDestID)
			}
			FwdRuleMap[int64MapKey][arrKey] = int64(intDestID)
		}
	}
}

func typeError(v interface{}) {
	logrus.WithField(fmt.Sprintf(`TypeOf(%v)`, v), reflect.TypeOf(v)).
		Fatal("Configuration file hasn't been set properly, type error: ")
}
