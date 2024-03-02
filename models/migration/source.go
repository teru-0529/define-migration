package migration

import (
	"fmt"
)

const sourceDir = "database/migrations"

// Source(Github)
type GithubOrganization struct {
	RequiredAuth bool   `mapstructure:"REQUIRED_AUTHENTICATION"`
	Owner        string `mapstructure:"GITHUB_OWNER"`
	User         string `mapstructure:"GITHUB_USER"`
	Token        string `mapstructure:"GITHUB_TOKEN"`
}

func FileSource(schema string) string {
	return fmt.Sprintf("file://%s/%s", schema, sourceDir)
}

// func (github Github) Source(schema Schema) string {

// }
