package models

type Config struct {
	DB struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
	SERVER struct {
		Host string `yaml:"host"`
	}
	DEBUG struct {
		Switch bool `yaml:"switch"`
		Login  bool `yaml:"login"`
		Cors   bool `yaml:"cors"`
	}
	Bot struct {
		Token      string `yaml:"token"`
		ScriptPath string `yaml:"scriptPath"`
		GetUpdate  string `yaml:"getUpdate"`
	}
}
