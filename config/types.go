package config

type Config struct {
	Application struct {
		ServiceName string
		ServerPort  string
	}
	UserDB DBConfig
	Token  TokenConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type TokenConfig struct {
	Key                string
	Expiry             int64
	RefreshTokenExpiry int64
}
