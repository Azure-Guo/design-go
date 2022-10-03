package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func Test(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./configs/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("err")
	}
	getString := viper.GetString("suggest.handler")
	fmt.Println(getString)
}
