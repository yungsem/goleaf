package config

import (
	_ "embed"
)

//go:embed env.yml
var EnvContent []byte

//go:embed local.yml
var LocalContent []byte

//go:embed prod.yml
var ProdContent []byte
