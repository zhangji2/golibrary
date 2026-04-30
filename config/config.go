package config

type AppConfigStruct struct {
	App struct {
		Name  string `yaml:"name"`
		Debug bool   `yaml:"debug"`
	} `yaml:"app"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
	Mysql struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"mysql"`
	RabbitMQ struct {
		Uri string `yaml:"uri"`
	} `yaml:"rabbitmq"`
}
