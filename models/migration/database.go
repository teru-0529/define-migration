package migration

// DB(Postgres)
type Postgres struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	Db       string `mapstructure:"POSTGRES_DB"`
}
