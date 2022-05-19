package storeprovider

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

type PostgresProvider struct {
	db        *sql.DB
	tablename string
}

func (pp *PostgresProvider) bootStrap(dbname string) error {

	fmt.Printf("DB Name: %s", dbname)
	dropDB := fmt.Sprintf(`DROP DATABASE IF EXISTS %s;`, dbname)
	createDB := fmt.Sprintf(`CREATE DATABASE %s;`, dbname)

	_, err := pp.db.Exec(dropDB)
	if err != nil {
		return err
	}

	_, err = pp.db.Exec(createDB)
	if err != nil {
		return err
	}

	return nil

}
func NewPostgresProvider(username string, password string, db_host string, dbname string, tablename string) (*PostgresProvider, error) {

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/?sslmode=disable", username, password, db_host)
	usageConnStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, db_host, dbname)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	provider := PostgresProvider{db: db, tablename: tablename}
	err = provider.bootStrap(dbname)
	if err != nil {
		return nil, err
	}
	err = db.Close()
	if err != nil {
		return nil, err
	}
	db, err = sql.Open("postgres", usageConnStr)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (rate float);", tablename))

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (0.25);", tablename))
	if err != nil {
		return nil, err
	}
	provider.db = db

	return &provider, nil

}

func (pp *PostgresProvider) Set(amount float32, rate float32) {

	queryStm := fmt.Sprintf("UPDATE rates SET amount = $1 WHERE rate = (select rate from %s limit 1);", pp.tablename)
	_, err := pp.db.Exec(queryStm, amount)
	if err != nil {

		// should be better error handling ... maybe the provider interface should expose error
		panic(err)
	}
}

func (pp *PostgresProvider) Get(amount float32) float32 {
	var rate float32
	queryStm := fmt.Sprintf("SELECT rate from %s limit 1", pp.tablename)
	row := pp.db.QueryRow(queryStm)

	switch err := row.Scan(&rate); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		return rate
	default:
		// should be better error handling ... maybe the provider interface should expose error
		panic(err)
	}

	// when no rate is found use 1 as multiplier which means no discount
	return 1
}
