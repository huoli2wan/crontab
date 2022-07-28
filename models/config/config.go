package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

//配置
type Config struct {
	ApiPort               int      `yaml:"apiPort"`
	ApiReadTimeout        int      `yaml:"apiReadTimeout"`
	ApiWriteTimeout       int      `yaml:"apiWriteTimeout"`
	EtcdEndpoints         []string `yaml:"etcdEndpoints"`
	EtcdDialTimeout       int      `yaml:"etcdDialTimeout"`
	MongodbUri            string   `yaml:"mongodbUri"`
	MongodbConnectTimeout int      `yaml:"mongodbConnectTimeout"`
}

var (
	G_config *Config
	filename = "config.yaml"
)

//InitConfig 初始化配置文件
func InitConfig() (err error) {
	var (
		content []byte
	)
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = yaml.Unmarshal(content, &G_config); err != nil {
		return
	}
	return
}
