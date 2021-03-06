package datamigration

import (
	"database/sql"
	"net/url"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres" // make driver available
	_ "github.com/mattes/migrate/source/file"     // make source available
	"github.com/pkg/errors"

	"github.com/chef/automate/lib/logger"
)

// NOTE (TC): These are IAM V2 specific data migrations. DO NOT APPLY ANY SCHEMA CHANGES HERE!
// Read more in v2_data_migrations.md.
//
// Config holds the information needed to connect to the database (PGURL), to
// find the migration SQL files (Path), and log debug messages (Logger).
type Config struct {
	PGURL  *url.URL
	Path   string
	Logger logger.Logger
}

type migrationLog struct {
	logger.Logger
}

func (migrationLog) Verbose() bool {
	return false
}

// Migrate executes all migrations we have
func (c *Config) Migrate() error {
	// m, err := migrate.New(addScheme(c.Path), c.PGURL.String())
	// if err != nil {
	// 	return errors.Wrap(err, "init v2 data migration")
	// }
	db, err := sql.Open("postgres", c.PGURL.String())
	if err != nil {
		return errors.Wrap(err, "init v2 sql client")
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "data_migrations",
	})
	if err != nil {
		return errors.Wrap(err, "init v2 data migration driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		addScheme(c.Path),
		"postgres", driver)
	if err != nil {
		return errors.Wrap(err, "init v2 data migration")
	}

	m.Log = migrationLog{c.Logger} // nolint: govet

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "execute v2 data migrations")
	}

	// The first error is trying to Close() the source. For our file source,
	// that's always nil
	_, err = m.Close() // nolint: gas
	return errors.Wrap(err, "close v2 data migrations connection")
}

func addScheme(p string) string {
	u := url.URL{}
	u.Scheme = "file"
	u.Path = p
	return u.String()
}
