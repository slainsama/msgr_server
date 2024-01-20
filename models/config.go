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
		APIUrl    string `yaml:"apiUrl"`
		Token     string `yaml:"token"`
		GetUpdate string `yaml:"getUpdate"`
		Methods   struct {
			SendMessage string `yaml:"sendMessage"`
			GetUpdates  string `yaml:"getUpdates"`
		}
	}
}
