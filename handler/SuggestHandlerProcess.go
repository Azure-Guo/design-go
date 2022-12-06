package handler

import (
	"design-go/pojo"
	"github.com/spf13/viper"
	"log"
	"reflect"
	"strings"
)

func ProcessHandler(userInfo pojo.UserInfo) []string {
	config := readConfigByKey("suggest.handler")
	split := strings.Split(config, ",")
	temp := make([]string, 0)
	for i, _ := range split {
		getType, ok := ConfigTypeMap[split[i]]
		if ok {
			ref := reflect.New(getType)
			handler := ref.Interface().(AbstractSuggestRequirementHandler)
			processHandler := handler.processHandler(userInfo, temp)
			temp = processHandler
		}
	}
	return temp
}

func readConfigByKey(key string) string {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./configs/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("viper.ReadInConfig err: ", err)
		return ""
	}
	return viper.GetString(key)
}
