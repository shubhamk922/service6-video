package migrate

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/ardanlabs/darwin/v3"
	"github.com/ardanlabs/darwin/v3/dialects/postgres"
	"github.com/ardanlabs/darwin/v3/drivers/generic"

	"example.com/service/business/data/sqldb"
	"github.com/jmoiron/sqlx"
)

var (
	//go:embed sql/migrate.sql
	migrateDoc string

	//go:embed sql/seed.sql
	seedDoc string
)

// Migrate attempts to bring the database up to date with the migrations
// defined in this package.
func Migrate(ctx context.Context, db *sqlx.DB) error {
	if err := sqldb.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	driver, err := generic.New(db.DB, postgres.Dialect{})
	if err != nil {
		return fmt.Errorf("construct darwin driver: %w", err)
	}

	d := darwin.New(driver, darwin.ParseMigrations(migrateDoc))
	return d.Migrate()
}

// Seed runs the seed document defined in this package against db. The queries
// are run in a transaction and rolled back if any fail.
func Seed(ctx context.Context, db *sqlx.DB) (err error) {

	fmt.Println("Enter to start seed")
	if err := sqldb.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	fmt.Println("Enter to start seed#################1")

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	fmt.Println("Enter to start seed#################2")

	defer func() {
		if errTx := tx.Rollback(); errTx != nil {
			if errors.Is(errTx, sql.ErrTxDone) {
				return
			}

			err = fmt.Errorf("rollback: %w", errTx)
			return
		}
	}()

	fmt.Println("Enter to start seed#################3")

	if _, err := tx.Exec(seedDoc); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	fmt.Println("Enter to start seed#################4")

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}
