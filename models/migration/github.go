package migration

import "fmt"

// Source(Github)

type Github struct {
	RequiredAuth bool   `mapstructure:"REQUIRED_AUTHENTICATION"`
	Owner        string `mapstructure:"GITHUB_OWNER"`
	User         string `mapstructure:"GITHUB_USER"`
	Token        string `mapstructure:"GITHUB_TOKEN"`
}

func (github *Github) baseUrl() string {
	if github.RequiredAuth {
		return fmt.Sprintf("github://%s:%s@%s", github.User, github.Token, github.Owner)
	} else {
		return fmt.Sprintf("github://%s", github.Owner)
	}
}
