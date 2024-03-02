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
)

// Schema list
type SchemaSet struct {
	Schemas []Schema `yaml:"schemas"`
}

type Schema struct {
	Name       string     `yaml:"name"`
	SourceType SourceType `yaml:"migrateFrom"`
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
			return nil, errors.New("schema file error: `migrateFrom` must be one of the [`github`,`local`]. ")
		}

		if schema.SourceType == GITHUB && schema.GithubUrl == "" {
			return nil, errors.New("schema file error: `githubUrl` is required if `migrateFrom` is [`github`]. ")
		}
	}

	// // INFO: originの設定
	// for i := range savedata.DeliveElements {
	// 	element := &savedata.DeliveElements[i]
	// 	original, err := savedata.getElement(element.Origin)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	element.Ref = original
	// }
	// for i := range savedata.Segments {
	// 	element := &savedata.Segments[i]
	// 	original, err := savedata.getElement(element.Key)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	element.Ref = original
	// }
	return &schemas, nil
}
