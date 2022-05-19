package conf

type Database struct {
	Name     string
	Driver   string
	Port     uint16
	Host     string
	User     string
	Password string
}

type AllowedHost struct {
	Host  string
	Https bool
}

type Server struct {
	ListenAddress           string
	AllowedHosts            []AllowedHost
	UseForwardedProtoHeader bool
	TemplatePath            string
	StaticPath              string
	PluginPath              string
}

type Config struct {
	Database Database
	Server   Server
}

func (c *Config) Check() error {
	return nil
}
