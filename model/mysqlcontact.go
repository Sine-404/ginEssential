package model

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Pwd      string `yaml:"pwd"`
}
