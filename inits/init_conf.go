package inits

import (
	"github.com/yungsem/goleaf/config"
	"github.com/yungsem/gotool/configuration"
)

var Conf *configuration.Conf

type configLoader struct {
}

func (c *configLoader) Load() map[string][]byte {
	m := make(map[string][]byte)
	m[configuration.ContentMapKeyEnv] = config.EnvContent
	m[configuration.ContentMapKeyLocal] = config.LocalContent
	m[configuration.ContentMapKeyProd] = config.ProdContent
	return m
}

// initConf 初始化配置
func initConf() {
	var loader configLoader
	Conf = configuration.LoadConfig(&loader)
}
