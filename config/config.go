package config

type ServerConfigurations struct {
	Port    string
	RunMode string
}

type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

type JWTConfiguration struct {
	SecretKey string
}

type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	JWT      JWTConfiguration
}
