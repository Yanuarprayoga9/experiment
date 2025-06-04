package services

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/rs/zerolog/log"
	"github.com/golang-migrate/migrate/v4"
	mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

)

func InitDatabase() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Format: username:password@tcp(host:port)/dbname
		dsn = "root:@tcp(127.0.0.1:3306)/bloggingdb"
	}

	// 1. Open SQL connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	// 2. Ping DB
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to ping database")
	}

	// 3. Run migrations using golang-migrate
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create migrate driver instance")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations", // migration files folder
		"bloggingdb",           // database name (as string, no need to match driver name)
		driver,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize migrate")
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal().Err(err).Msg("failed to run database migration")
	} else {
		log.Info().Msg("database migrated successfully")
	}

	// // 4. Custom logic (opsional)
	// _, errCheck := db.Exec(sqlquery.CASBIN_CREATE)
	// if errCheck != nil {
	// 	log.Fatal().Err(errCheck).Msg("failed to check if table exists")
	// }

	// _, errCheck = db.Exec(sqlquery.CASBIN_VIEW)
	// if errCheck != nil {
	// 	log.Fatal().Err(errCheck).Msg("failed to check if view exists")
	// }

	// _, errCheck = db.Exec(sqlquery.CASBIN_INIT)
	// if errCheck != nil {
	// 	log.Fatal().Err(errCheck).Msg("failed to initialize database")
	// }

	return db
}
