package migrations

import (
	"fmt"
	"log"

	"github.com/motephyr/longapp/app"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	n, err := migrate.Exec(app.Http.Database.DB, app.Http.Database.Default.Driver, migrations, migrate.Up)
	if err != nil {
		log.Println(err)
		panic(err)
		// Handle errors!
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
