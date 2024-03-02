package migration

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

// migration source
type SourceType string

var (
	GITHUB     SourceType = "github"
	LOCAL_FILE SourceType = "local"
	DEFAULT    SourceType = "settingFile"
)

// Schema list
type SchemaSet struct {
	Schemas []Schema `yaml:"schemas"`
	sMap    map[string]Schema
}

type Schema struct {
	Name       string     `yaml:"name"`
	SourceType SourceType `yaml:"sourceType"`
	GithubUrl  string     `yaml:"githubUrl"`
	GithubTag  string     `yaml:"tag"`
}

func NewSchemaSet(path string) (*SchemaSet, error) {
	// INFO: read
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %s", err.Error())
	}

	// INFO: unmarchal
	var schemas SchemaSet
	err = yaml.Unmarshal([]byte(file), &schemas)
	if err != nil {
		return nil, err
	}

	// INFO: validation
	for _, schema := range schemas.Schemas {
		if !slices.Contains([]SourceType{GITHUB, LOCAL_FILE}, schema.SourceType) {
			return nil, errors.New("schema file error: 'sourceType' must be one of the [\"github\",\"local\"]. ")
		}

		if schema.SourceType == GITHUB && schema.GithubUrl == "" {
			return nil, errors.New("schema file error: 'githubUrl' is required if 'sourceType' is [\"github\"]. ")
		}
	}

	// INFO: mapの設定
	schemas.sMap = map[string]Schema{}
	for _, schema := range schemas.Schemas {
		schemas.sMap[schema.Name] = schema
	}

	return &schemas, nil
}

// 存在チェック
func (schemas *SchemaSet) Exist(name string) bool {
	_, ok := schemas.sMap[name]
	return ok
}
