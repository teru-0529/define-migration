package migration

import (
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
)

// ----+----+----+----+----+----+----+----+----+----

// Source list
type SourceSet struct {
	SourceArray []Source `yaml:"sources"`
	sourceMap   map[string]Source
	github      Github
}

type Source struct {
	SchemaName string     `yaml:"schema"`
	Repository string     `yaml:"gitRepository"`
	Tag        string     `yaml:"gitTag"`
	SourceType SourceType `yaml:"sourceType"`
}

// ソースファイルの読み込み
func NewSourceSet(path string, github Github, useLocal bool) (*SourceSet, error) {
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

	// INFO: sourceTypeの設定
	for i, source := range sources.SourceArray {
		// 強制ローカルファイル利用の場合
		if useLocal {
			sources.SourceArray[i].SourceType = LOCAL_FILE
		}

		// ソース種類が不明文字列の場合
		if !slices.Contains([]SourceType{GITHUB, LOCAL_FILE}, source.SourceType) {
			fmt.Printf(
				"warnings: 'sourceType'[\"%s\"] of 'schema'[\"%s\"] is converted \"local\" , because of 'sourceType' must be one of the [\"github\",\"local\"]\n",
				source.SourceType,
				source.SchemaName,
			)
			sources.SourceArray[i].SourceType = LOCAL_FILE
		}

		// リポジトリ設定がないのに、ソース種類がGithubの場合
		if source.SourceType == GITHUB && source.Repository == "" {
			fmt.Printf(
				"warnings: 'sourceType' of 'schema'[\"%s\"] is converted \"local\" , because of required 'githubRepository' if 'sourceType' is [\"github\"]\n",
				source.SchemaName,
			)
			sources.SourceArray[i].SourceType = LOCAL_FILE
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

// INFO: ソース存在確認
func (sources *SourceSet) Exist(schemaName string) bool {
	_, isOk := sources.sourceMap[schemaName]
	return isOk
}

// INFO: ソース入力先
func (sources *SourceSet) SourceUrl(schemaName string) string {
	source, existSource := sources.sourceMap[schemaName]
	if !existSource {
		panic("source not exist!")
	}

	if source.SourceType == GITHUB {
		return source.githubSource(sources.github)
	} else {
		return source.fileSource()
	}
}

// ----+----+----+----+----+----+----+----+----+----

// githubソース
func (source *Source) githubSource(github Github) string {
	if source.Tag == "" {
		return fmt.Sprintf("%s/%s/%s", github.baseUrl(), source.Repository, sourceDir)
	} else {
		return fmt.Sprintf("%s/%s/%s#%s", github.baseUrl(), source.Repository, sourceDir, source.Tag)
	}
}

// fileソース
func (source *Source) fileSource() string {
	return fmt.Sprintf("file://./%s/%s", source.SchemaName, sourceDir)
}
