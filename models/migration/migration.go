package migration

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
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
func MigrateUp(io MigrationIO) error {
	fmt.Println("-----", io.Schema, "migrate-up", "-----")
	io.Show()
	err := getMigration(io)
	if err != nil {
		return err
	}
	fmt.Println(io.Schema, "up")
	return nil
}

// INFO: DOWN
func MigrateDown(io MigrationIO) error {
	fmt.Println("-----", io.Schema, "migrate-down", "-----")
	io.Show()
	err := getMigration(io)
	if err != nil {
		return err
	}
	fmt.Println(io.Schema, "down")
	return nil
}

func getMigration(io MigrationIO) error {
	m, err := migrate.New(io.Source, io.Database)
	if err != nil {
		return err
	}
	version, dirty, err := m.Version()
	fmt.Println(version, dirty, err)
	return nil
}
