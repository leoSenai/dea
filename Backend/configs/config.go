package configs

import "github.com/spf13/viper"

var cfg *Config
var vi *viper.Viper = viper.New()

func init() {

	vi.SetDefault("api.port", "9000")
	vi.SetDefault("database.host", "localhost")
	vi.SetDefault("database.port", "3306")
	vi.SetDefault("database.user", "root")
	vi.SetDefault("database.pass", "root")
}

func Load() error {

	vi.SetConfigType("toml")
	vi.AddConfigPath(".")
	vi.SetConfigName("config")
	vi.SetConfigFile("config.toml")

	err := vi.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(Config)

	cfg.API = APIConfig{
		Port: vi.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     vi.GetString("database.host"),
		Port:     vi.GetString("database.port"),
		User:     vi.GetString("database.user"),
		Pass:     vi.GetString("database.pass"),
		Database: vi.GetString("database.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
