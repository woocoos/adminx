//go:build ignore

package main

import (
	atlas "ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/woocoos/adminx/ent/migrate"
	"log"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "mysql://root@localhost:3306/portal", "database dsn:mysql://root:pass@localhost:3306/test")
	name = flag.String("name", "", "migration name")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithForeignKeys(false),
	}
	if *name == "" {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, *dsn, *name, opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
