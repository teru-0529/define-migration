package migration

import "fmt"

// DB(Postgres)
type Postgres struct {
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Host     string `mapstructure:"POSTGRES_HOST_NAME"`
	Port     string `mapstructure:"POSTGRES_HOST_PORT"`
	Db       string `mapstructure:"POSTGRES_DB"`
}

// INFO: データベース出力先
func (postgres *Postgres) DatabaseUrl(schema string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&x-migrations-table=migrate_%s",
		postgres.User,
		postgres.Password,
		postgres.Host,
		postgres.Port,
		postgres.Db,
		schema,
	)
}
