package database

import (
	"database/sql"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/kulado/sqlxmigrate"
	"github.com/ndious/snacked/internal"
)

type MigrationSchema struct {
	Up   string
	Down string
}

type MigrationFileSchema struct {
	Migration MigrationSchema `toml:"migration"`
}

func Migrate() ([]string, error) {

	var migrations []*sqlxmigrate.Migration
	var filesPath []string
	var err error

	db := GetDb()

	for _, file := range getMigrationFiles() {
		var change MigrationFileSchema

		content, _ := os.ReadFile(migrationFilePath(file))
		filesPath = append(filesPath, file.Name())

		if err := toml.Unmarshal(content, &change); err != nil {
			return nil, err 
		}

		migrationID := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		migrations = append(migrations, &sqlxmigrate.Migration{
			ID: migrationID,
			Migrate: func(tx *sql.Tx) error {
				_, err = tx.Exec(change.Migration.Up)
				return err
			},
			Rollback: func(tx *sql.Tx) error {
				_, err = tx.Exec(change.Migration.Down)
				return err
			},
		})
	}

	m := sqlxmigrate.New(db, sqlxmigrate.DefaultOptions, migrations)

	defer db.Close()

	if err = m.Migrate(); err != nil {
		return nil, err
	}

	return filesPath, nil
}

func migrationFilePath(file fs.DirEntry) string {
	return filepath.Join(migrationPath(), file.Name())
}

func migrationPath() string {
	return internal.GetDir("migrations")
}

func getMigrationFiles() []fs.DirEntry {
	files, err := os.ReadDir(migrationPath())

	if err != nil {
		panic(err)
	}

	return files
}
