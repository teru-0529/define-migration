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
	fmt.Println("")
	fmt.Println("-----", io.Schema, "migrate-up", "-----")
	io.Show()
	// マイグレーションの取得
	mig, err := getMigration(io)
	if err != nil {
		return err
	}
	// 実行前バージョン
	showVersion(mig, "before")
	// 実行
	if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	// 実行後バージョン
	showVersion(mig, "after ")

	return nil
}

// INFO: DOWN
func MigrateDown(io MigrationIO) error {
	fmt.Println("")
	fmt.Println("-----", io.Schema, "migrate-down", "-----")
	io.Show()
	// マイグレーションの取得
	mig, err := getMigration(io)
	if err != nil {
		return err
	}
	// 実行前バージョン
	showVersion(mig, "before")
	// 実行
	if err := mig.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	// 実行後バージョン
	showVersion(mig, "after ")

	return nil
}

func getMigration(io MigrationIO) (*migrate.Migrate, error) {
	mig, err := migrate.New(io.Source, io.Database)
	if err != nil {
		return nil, err
	}
	return mig, nil

	// version, dirty, _ := m.Version()
	// // if err != nil && err != migrate.ErrInvalidVersion {
	// // 	fmt.Println(err)
	// // 	fmt.Println(130)
	// // 	return nil, err
	// // }
	// if dirty {
	// 	fmt.Println("force execute current version sql")
	// 	m.Force(int(version))
	// }

	// fmt.Printf("before version: %d\n", version)
	// return m, nil
}

func showVersion(mig *migrate.Migrate, header string) {
	version, dirty, err := mig.Version()
	if dirty {
		fmt.Println("force execute current version sql")
		mig.Force(int(version))
	}
	if err != nil {
		fmt.Printf("%s version: %d, %v\n", header, version, err)

	} else {
		fmt.Printf("%s version: %d\n", header, version)

	}

	// if err != nil {
	// 	return err
	// }
}
