package config

import (
	_ "embed"
	"gopkg.in/yaml.v2"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

//go:embed env.yml
var envContent []byte

//go:embed local.yml
var localContent []byte

//go:embed prod.yml
var prodContent []byte

// loadEnv 解析 env.yml 文件的内容，放到 env 中
func loadEnv() *env {
	// 定义一个 env
	e := env{}

	// 解析 yml
	err := yaml.Unmarshal(envContent, &e)
	if err != nil {
		panic("配置文件env加载失败")
	}

	return &e
}

// LoadConfig 解析配置文件的内容
func LoadConfig() *Conf {
	// 加载 env
	e := loadEnv()

	var content []byte
	if e.Env == envLocal {
		content = localContent
	} else if e.Env == envProd {
		content = prodContent
	} else {
		// 预留
	}

	// 定义一个 conf
	c := Conf{}

	// 解析 yml
	err := yaml.Unmarshal(content, &c)
	if err != nil {
		panic("配置文件内容加载失败")
	}

	return &c
}
