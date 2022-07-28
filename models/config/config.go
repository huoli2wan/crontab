package config

import (
	"encoding/json"
	"io/ioutil"
)

//配置
type Config struct {
	ApiPort               int      `yaml:"apiPort"`
	ApiReadTimeout        int      `yaml:"apiReadTimeout"`
	ApiWriteTimeout       int      `yaml:"apiWriteTimeout"`
	EtcdEndpoints         []string `yaml:"etcdEndpoints"`
	EtcdDialTimeout       int      `yaml:"etcdDialTimeout"`
	WebRoot               string   `yaml:"webroot"`
	MongodbUri            string   `yaml:"mongodbUri"`
	MongodbConnectTimeout int      `yaml:"mongodbConnectTimeout"`
}

var (
	G_config *Config
	filename = "config.yaml"
)

func InitConfig() (err error) {
	var (
		content []byte
	)
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(content, &G_config); err != nil {
		return
	}
	return
}
