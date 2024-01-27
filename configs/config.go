package config

type MongoDBConfig struct {
    URI      string `yaml:"uri"`
    Database string `yaml:"database"`
    Timeout  int    `yaml:"timeout"`
}

type ServerConfig struct {
    Host string `yaml:"host"`
    Port int    `yaml:"port"`
}