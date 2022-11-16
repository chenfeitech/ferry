package config

import "github.com/spf13/viper"

type Database struct {
	Dbtype   string
	Host     string
	Port     int
	Name string
	Username string
	Password string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port:     cfg.GetInt("port"),
		Dbtype:   cfg.GetString("dbType"),
		Host:     cfg.GetString("host"),
		Name: cfg.GetString("name"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
	}
}

var DatabaseConfig = new(Database)

type Miniprogram struct {
	Appid		string
	Secret		string
	Redisaddr	string
}

func InitMiniprogram(cfg *viper.Viper) *Miniprogram {
	return &Miniprogram{
		Appid:     cfg.GetString("appid"),
		Secret:   cfg.GetString("secret"),
		Redisaddr:     cfg.GetString("redisaddr"),
	}
}

var MiniprogramConfig = new(Miniprogram)