/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teru-0529/define-migration/models/migration"
)

// バージョン
var (
	version     string
	releaseDate string
)

// FLAG
var envFile string
var sourceFile string
var sourceTypeStr string

var (
	postgres migration.Postgres
	github   migration.Github
	sources  migration.SourceSet
)

// ソース種類（Flag入力）
var sourceType migration.SourceType

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "define-migration",
	Short: "service for table migration.",
	Long:  "service for table migration.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ver string, date string) {
	version = ver
	releaseDate = date

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(mirrorCmd)
	rootCmd.AddCommand(versionCmd)

	rootCmd.PersistentFlags().StringVarP(&envFile, "env-file", "", ".env", "env file")
	rootCmd.PersistentFlags().StringVarP(&sourceFile, "source-file", "", "source-setting.yaml", "source setting file")

	rootCmd.PersistentFlags().StringVarP(&sourceTypeStr, "sourceType", "T", "sourceFile", "if setting one of the [\"github\",\"local\"], force change sourceType.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("env")
	viper.SetConfigFile(envFile)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.Unmarshal(&postgres); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&github); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// ソースファイルの読込み
	sources_, err := migration.NewSourceSet(sourceFile, github)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	sources = *sources_
	fmt.Printf("Using source file: %s\n", sourceFile)

	// ソース種類の変換
	sourceType = migration.SourceType(sourceTypeStr)
	if !sourceType.Varidate() {
		log.Println("input flag error: 'sourceType' must be one of the [\"github\",\"local\"]. ")
		os.Exit(1)
	}
}
