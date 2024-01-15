package migrator

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "Storage path")
	flag.StringVar(&migrationsPath, "migrations-path", "", "Migrations path")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "Migrations table")
	flag.Parse()

	if storagePath == "" || migrationsPath == "" {
		panic("storage-path and migrations-path are required")
	}

	m, err := migrate.New("file://"+migrationsPath, fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable))
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("Migrations applied")
}
