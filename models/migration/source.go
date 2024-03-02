package migration

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

const sourceDir = "database/migrations"

// ----+----+----+----+----+----+----+----+----+----
// migration source
type SourceType string

var (
	GITHUB     SourceType = "github"
	LOCAL_FILE SourceType = "local"
	DEFAULT    SourceType = "sourceFile"
)

// SourceTypeの種類チェック
func (sourceType SourceType) Varidate() bool {
	return slices.Contains([]SourceType{GITHUB, LOCAL_FILE, DEFAULT}, sourceType)
}

// ----+----+----+----+----+----+----+----+----+----

// Source list
type SourceSet struct {
	SourceArray []Source `yaml:"sources"`
	sourceMap   map[string]Source
	github      Github
}

type Source struct {
	SchemaName string     `yaml:"schema"`
	SourceType SourceType `yaml:"sourceType"`
	Repository string     `yaml:"gitRepository"`
	Tag        string     `yaml:"gitTag"`
}

func NewSourceSet(path string, github Github) (*SourceSet, error) {
	// INFO: read
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %s", err.Error())
	}

	// INFO: unmarchal
	var sources SourceSet
	err = yaml.Unmarshal([]byte(file), &sources)
	if err != nil {
		return nil, err
	}

	// INFO: validation
	for _, source := range sources.SourceArray {
		if !slices.Contains([]SourceType{GITHUB, LOCAL_FILE}, source.SourceType) {
			return nil, errors.New("schema file error: 'sourceType' must be one of the [\"github\",\"local\"]. ")
		}

		if source.SourceType == GITHUB && source.Repository == "" {
			return nil, errors.New("schema file error: 'githubRepository' is required if 'sourceType' is [\"github\"]. ")
		}
	}

	// INFO: mapの設定
	sources.sourceMap = map[string]Source{}
	for _, source := range sources.SourceArray {
		sources.sourceMap[source.SchemaName] = source
	}

	// INFO: githubの設定
	sources.github = github

	return &sources, nil
}

// スキーマ存在チェック
func (sources *SourceSet) Exist(schemaName string) bool {
	_, isOk := sources.sourceMap[schemaName]
	return isOk
}

// ----+----+----+----+----+----+----+----+----+----

// Source(Github)

func FileSource(schema string) string {
	return fmt.Sprintf("file://%s/%s", schema, sourceDir)
}

// func (github Github) Source(schema Schema) string {

// }
