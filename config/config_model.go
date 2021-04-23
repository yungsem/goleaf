package config

type env struct {
	Env string `yaml:"env"`
}

type Conf struct {
	Server struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	}
	Log struct {
		Output string `yaml:"output"`
		Level  string `yaml:"level"`
		Path   string `yaml:"path"`
	}
	Mysql struct {
		Url      string `yaml:"url"`
		DB       string `yaml:"db"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}
