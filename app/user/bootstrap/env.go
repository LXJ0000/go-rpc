package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv string `yaml:"appEnv"`

	Server   Server             `yaml:"server"`
	MySQL    MySQL              `yaml:"mysql"`
	Redis    Redis              `yaml:"redis"`
	Etcd     Etcd               `yaml:"etcd"`
	Services map[string]Service `yaml:"services"`
	Domain   map[string]Domain  `yaml:"domain"`
}

type Server struct {
	Port               string `yaml:"port"`
	Version            string `yaml:"version"`
	JwtSecret          string `yaml:"jwtSecret"`
	SnowFlakeStartTime string `yaml:"snowFlakeStartTime"`
	SnowFlakeMachine   int64  `yaml:"snowFlakeMachine"`
	ContextTimeout     int    `yaml:"contextTimeout"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	DSN string `yaml:"dsn"`
}

type Etcd struct {
	Address string `yaml:"address"`
}

type Service struct {
	Name        string `yaml:"name"`
	LoadBalance bool   `yaml:"loadBalance"`
	Addr        string `yaml:"addr"`
}

type Domain struct {
	Name string `yaml:"name"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile("config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .yml : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
