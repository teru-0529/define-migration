package migration

import (
	"fmt"
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	// _ "github.com/golang-migrate/migrate/v4/source/github"
)

type MigrationIO struct {
	Schema   string
	Source   string
	Database string
}

// INFO: IO情報の表示
func (io *MigrationIO) Show() {
	fmt.Printf("[source  ] %s\n", io.Source)
	fmt.Printf("[database] %s\n", io.Database)
}

// INFO: UP
func MigrateUp(migrationIo MigrationIO) error {
	fmt.Println("-----", migrationIo.Schema, "migrate-up", "-----")
	migrationIo.Show()
	fmt.Println(migrationIo.Schema, "up")
	return nil
}

// INFO: DOWN
func MigrateDown(migrationIo MigrationIO) error {
	fmt.Println("-----", migrationIo.Schema, "migrate-down", "-----")
	migrationIo.Show()
	fmt.Println(migrationIo.Schema, "down")
	return nil
}

// func getMigration(schemaName string) {
// 	source:=
// }
