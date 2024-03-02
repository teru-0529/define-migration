package migration

// Source(Github)
type Github struct {
	User  string `mapstructure:"GITHUB_USER"`
	Token string `mapstructure:"GITHUB_TOKEN"`
}
