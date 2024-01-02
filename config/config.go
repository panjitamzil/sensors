package config

type Config struct {
	Application Application
	Database    DatabaseConfig
	GRPC        GRPCConfig
}

type Application struct {
	ServiceName string
	Port        string
}

type DatabaseConfig struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Name                  string
	MaxIdleConnection     int
	MaxOpenConnection     int
	MaxLifetimeConnection int
	MaxIdletimeConnection int
}

type GRPCConfig struct {
	ServiceName string
	Host        string
	Port        string
}
