package core

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type configS struct {
	Protocol string
	Host     string
	Mysql    string
	Redis    string
	Gkey     string
	Gsecret  string
}

var Config configS

func initConfig() {
	var err error
	configData, err := ioutil.ReadFile("./conf/default.yaml")
	if err != nil {
		log.Fatalln("[-] Read config failed.", err)
	}
	err = yaml.Unmarshal(configData, &Config)
	if err != nil {
		log.Fatalln("[-] Unmarshal config data error.", err)
	}
	fmt.Println(Config.Gsecret)
}
